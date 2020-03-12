package main

import (
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func main() {
	createSecret("pkb", "default")

}
func createSecret(name, namespace string) {

	configMapKeyVal := "dummyKey"
	nonceVal := "dummyNonce"

	data := map[string][]byte{
		"configMapKey": []byte(configMapKeyVal),
		"nonce":        []byte(nonceVal),
	}
	label := map[string]string{
		"comment": "This is dummy secret",
	}

	secret := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name, Labels: label},
		Data:       data,
	}

	e := json.NewYAMLSerializer(json.DefaultMetaFactory, nil, nil)
	err := e.Encode(&secret, os.Stdout)
	if err != nil {
		panic(err)
	}

}
