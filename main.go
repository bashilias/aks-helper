package main

// Copyright 2022 Ilias Yacoubi (hi@ilias.sh)

// Goal of this application is to learn and get familiar with the client-go package.

import (
	"flag"
	"fmt"
)

<<<<<<< HEAD
func getCluster() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// return clientset
	return kubernetes.NewForConfig(config)

}

// input clientset. return ingress items.
func getIngress(clientset kubernetes.Clientset) ([]v1.Ingress, error) {

	// get all ingresses
	ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// return ingress items
	return ingresses.Items, nil

}

func inspectIngress(i []v1.Ingress) ([]string, []string, []bool, []bool) {

	// slice for hosts
	var hs []string
	// slice for backend
	var bs []string
	// slice for annotation keys
	var ls []string
	// slice for whitelist
	var wl []bool
	// slice for helm annotation
	var hl []bool

	for value := range i {

		ingRuleHost := &i[value].Spec.Rules[0].Host
		ingRulePath := &i[value].Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path
		ingBackendService := &i[value].Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName
		ingAnnotation := &i[value].Annotations

		// use regexp to replace these characters with nothing
		re, err := regexp.Compile(`[().*?$+]`)
		if err != nil {
			log.Fatal(err)
		}

		*ingRulePath = re.ReplaceAllString(*ingRulePath, "")

		// look for ' | ' in path, split it and put value in a slice.
		split := strings.Split(*ingRulePath, "|")
		for _, value := range split {
			// check is value doesn't start with '/', if not add '/'.
			if !strings.HasPrefix(value, "/") {
				value = "/" + value

			}
			fullSlug := *ingRuleHost + value
			hs = append(hs, fullSlug)
			bs = append(bs, *ingBackendService)

		}

		// add key of maps into slice 'ls'
		for i, _ := range *ingAnnotation {
			ls = append(ls, i)

		}

		for _, j := range ls {
			// Check if nginx whitelist annotation is there.
			if j == "nginx.ingress.kubernetes.io/whitelist-source-range" {

				wl = append(wl, true) //possible whitelist
			} else {
				wl = append(wl, false) //no nginx whitelist
			}
			// Check if helm annotation is there.
			if j == "meta.helm.sh/release-name" {

				hl = append(wl, true) //possible helm chart
			} else {
				hl = append(wl, false) //no helm chart
			}

		}

	}
	return hs, bs, wl, hl

}

func statusChecker(s string) bool {
	_, err := http.Get(s)
	var resp bool

	if err != nil {
		resp = false
	} else {
		resp = true

	}
	return resp
=======
var (
	ing bool
)
>>>>>>> 02ee15f (Flags added & Ingress code moved to ingress.go)

func init() {
	flag.BoolVar(&ing, "ing", false, "Check cluster for dead Ingresses.")
}

func main() {
<<<<<<< HEAD
	clientset, _ := getCluster()
	ingItems, _ := getIngress(*clientset)
	hs, _, wl, hl := inspectIngress(ingItems)

	i := 0

	for _, host := range hs {
		url := "http://" + host

		if !statusChecker(url) && !wl[i] && !hl[i] {
			fmt.Printf("🔴 %s \n", host)
			fmt.Printf("Remove Ingress. Y/N: ")
			var di string
			fmt.Scanln(&di)
			if di == "Y" || di == "y" {
				fmt.Println("Removing Ingress...")
			}

		}

		i++
=======

	flag.Parse()

	if ing {
		clientset, _ := getCluster()
		ingItems, _ := getIngress(*clientset)
		hs, _, wl, hl := inspectIngress(ingItems)
		i := 0
		for _, host := range hs {
			url := "http://" + host

			if !statusChecker(url) {
				fmt.Printf("🔴 %s \n\t Whitelist: %s \n\t Helm: %s\n", host, wl[i], hl[i])

			}

			i++

		}
		fmt.Printf("\n🔎💻 %d URL's \n", len(hs))
>>>>>>> 02ee15f (Flags added & Ingress code moved to ingress.go)

	}

	fmt.Printf("\n🔎💻 %d URL's \n", len(hs))

}
