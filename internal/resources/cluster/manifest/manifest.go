/*
Copyright (c) 2020 VMware, Inc. All rights reserved.

Proprietary and confidential.

Unauthorized copying or use of this file, in any medium or form,
is strictly prohibited.
*/

package manifest

import (
	"fmt"

	"github.com/pkg/errors"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"
)

func Create(
	k8sclient k8sClient.Client,
	k8sManifest []byte,
	forceClean bool,
) error {
	if k8sclient == nil {
		return errors.New("kubernetes client cannot be empty")
	}

	manifests, err := getManifests(string(k8sManifest))
	if err != nil {
		return errors.WithMessage(err, "failure to fetch attach manifests")
	}

	toBeCleaned, err := objectsToBeCleaned(k8sclient, manifests, forceClean)
	if err != nil && forceClean {
		return errors.WithMessage(err, "error while cleaning up the resources")
	}

	if len(toBeCleaned) != 0 {
		fmt.Println("Provided kubeconfig cannot be used to attach, reason:")

		for _, cleanup := range toBeCleaned {
			fmt.Println(cleanup)
		}

		return errors.New("please clean up the above mentioned k8s objects or follow cluster detach steps and retry")
	}

	err = createObjects(k8sclient, manifests)
	if err != nil {
		return errors.WithMessage(err, "error while attaching the cluster")
	}

	fmt.Println("TMC resources applied to the cluster successfully")

	return nil
}
