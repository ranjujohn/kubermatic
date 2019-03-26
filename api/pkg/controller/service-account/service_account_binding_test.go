package serviceaccount

import (
	"context"
	"testing"

	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/handler/test"

	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/client-go/kubernetes/scheme"

	controllerclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestCreateBinding(t *testing.T) {

	tests := []struct {
		name                      string
		saName                    string
		existingKubermaticObjects []runtime.Object
		expectedBinding           *kubermaticv1.UserProjectBinding
	}{
		{
			name:   "scenario 1: this test creates binding for service account",
			saName: "serviceaccount-abcd",
			existingKubermaticObjects: []runtime.Object{
				test.GenDefaultProject(),
				test.GenServiceAccount("abcd", "test", "editors", "my-first-project-ID"),
			},
			expectedBinding: func() *kubermaticv1.UserProjectBinding {
				binding := test.GenBinding("my-first-project-ID", "serviceaccount-abcd@sa.kubermatic.io", "editors")
				binding.OwnerReferences[0].Kind = kubermaticv1.UserKindName
				binding.OwnerReferences[0].Name = "serviceaccount-abcd"
				binding.Labels = map[string]string{kubermaticv1.ProjectIDLabelKey: "my-first-project-ID"}
				binding.Spec.Group = "editors-my-first-project-ID"
				return binding
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup the test scenario

			scheme := scheme.Scheme
			if err := kubermaticv1.AddToScheme(scheme); err != nil {
				t.Fatal(err)
			}
			kubermaticFakeClient := fake.NewFakeClientWithScheme(scheme, test.existingKubermaticObjects...)

			// act
			target := reconcileServiceAccountProjectBinding{ctx: context.TODO(), Client: kubermaticFakeClient}

			_, err := target.Reconcile(reconcile.Request{NamespacedName: types.NamespacedName{Name: test.saName}})

			// validate
			if err != nil {
				t.Fatal(err)
			}
			bindings := &kubermaticv1.UserProjectBindingList{}
			err = kubermaticFakeClient.List(context.TODO(), &controllerclient.ListOptions{}, bindings)
			if err != nil {
				t.Fatal(err)
			}

			if len(bindings.Items) != 1 {
				t.Fatalf("wrong number of bindigs, expectd 1 got %d", len(bindings.Items))
			}

			if !equality.Semantic.DeepEqual(bindings.Items[0].Labels, test.expectedBinding.Labels) {
				t.Fatalf("%v", diff.ObjectDiff(bindings.Items[0].Labels, test.expectedBinding.Labels))
			}

			if !equality.Semantic.DeepEqual(bindings.Items[0].OwnerReferences, test.expectedBinding.OwnerReferences) {
				t.Fatalf("%v", diff.ObjectDiff(bindings.Items[0].OwnerReferences, test.expectedBinding.OwnerReferences))
			}

			if !equality.Semantic.DeepEqual(bindings.Items[0].Spec, test.expectedBinding.Spec) {
				t.Fatalf("%v", diff.ObjectDiff(bindings.Items[0].Spec, test.expectedBinding.Spec))
			}

		})
	}
}
