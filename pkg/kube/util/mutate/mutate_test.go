package mutate_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	qstsv1a1 "code.cloudfoundry.org/quarks-statefulset/pkg/kube/apis/quarksstatefulset/v1alpha1"
	cfakes "code.cloudfoundry.org/quarks-statefulset/pkg/kube/controllers/fakes"
	"code.cloudfoundry.org/quarks-statefulset/pkg/kube/util/mutate"
	"code.cloudfoundry.org/quarks-utils/pkg/pointers"
)

var _ = Describe("Mutate", func() {
	var (
		ctx    context.Context
		client *cfakes.FakeClient
	)

	BeforeEach(func() {
		client = &cfakes.FakeClient{}
	})

	Describe("QuarksStatefulSetMutateFn", func() {
		var (
			eSts *qstsv1a1.QuarksStatefulSet
		)

		BeforeEach(func() {
			eSts = &qstsv1a1.QuarksStatefulSet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: qstsv1a1.QuarksStatefulSetSpec{
					Template: appsv1.StatefulSet{
						Spec: appsv1.StatefulSetSpec{
							Replicas: pointers.Int32(1),
						},
					},
				},
			}
		})

		Context("when the quarksStatefulSet is not found", func() {
			It("creates the quarksStatefulSet", func() {
				client.GetCalls(func(context context.Context, nn types.NamespacedName, object runtime.Object) error {
					return apierrors.NewNotFound(schema.GroupResource{}, nn.Name)
				})

				ops, err := controllerutil.CreateOrUpdate(ctx, client, eSts, mutate.QuarksStatefulSetMutateFn(eSts))
				Expect(err).ToNot(HaveOccurred())
				Expect(ops).To(Equal(controllerutil.OperationResultCreated))
			})
		})

		Context("when the quarksStatefulSet is found", func() {
			It("updates the quarksStatefulSet when spec is changed", func() {
				client.GetCalls(func(context context.Context, nn types.NamespacedName, object runtime.Object) error {
					switch object := object.(type) {
					case *qstsv1a1.QuarksStatefulSet:
						existing := &qstsv1a1.QuarksStatefulSet{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "foo",
								Namespace: "default",
							},
							Spec: qstsv1a1.QuarksStatefulSetSpec{
								Template: appsv1.StatefulSet{
									Spec: appsv1.StatefulSetSpec{
										Replicas: pointers.Int32(2),
									},
								},
							},
						}
						existing.DeepCopyInto(object)

						return nil
					}

					return apierrors.NewNotFound(schema.GroupResource{}, nn.Name)
				})
				ops, err := controllerutil.CreateOrUpdate(ctx, client, eSts, mutate.QuarksStatefulSetMutateFn(eSts))
				Expect(err).ToNot(HaveOccurred())
				Expect(ops).To(Equal(controllerutil.OperationResultUpdated))
			})

			It("does not update the quarksStatefulSet when nothing is changed", func() {
				client.GetCalls(func(context context.Context, nn types.NamespacedName, object runtime.Object) error {
					switch object.(type) {
					case *qstsv1a1.QuarksStatefulSet:
						return nil
					}

					return apierrors.NewNotFound(schema.GroupResource{}, nn.Name)
				})
				ops, err := controllerutil.CreateOrUpdate(ctx, client, eSts, mutate.QuarksStatefulSetMutateFn(eSts))
				Expect(err).ToNot(HaveOccurred())
				Expect(ops).To(Equal(controllerutil.OperationResultNone))
			})
		})
	})
})
