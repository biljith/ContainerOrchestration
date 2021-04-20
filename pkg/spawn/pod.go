package spawn

import (
	isv1alpha1 "dynamic-orchestration/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPod(vertex *isv1alpha1.Vertex) *corev1.Pod {
	labels := map[string]string{
		"app": vertex.Name,
	}

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vertex.Name,
			Namespace: vertex.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"printenv"},
				},
			},
			RestartPolicy: corev1.RestartPolicyOnFailure,
		},
	}
}
