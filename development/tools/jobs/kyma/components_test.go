package kyma

import (
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/releases"
	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/jobsuite"
)

var components = []struct {
	path              string
	image             string
	suite             func(config *jobsuite.Config) jobsuite.Suite
	additionalOptions []jobsuite.Option
}{
	{path: "api-controller", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "apiserver-proxy", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "application-broker", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "application-connectivity-certs-setup-job", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-connectivity-validator", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-gateway", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-operator", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-registry", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "asset-metadata-service", image: tester.ImageGolangBuildpack1_11},
	{path: "asset-store-controller-manager", image: tester.ImageGolangKubebuilder2BuildpackLatest},
	{path: "asset-upload-service", image: tester.ImageGolangBuildpack1_11},
	{path: "cms-controller-manager", image: tester.ImageGolangKubebuilder2BuildpackLatest},
	{path: "backup-plugins", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "cms-services", image: tester.ImageGolangBuildpack1_12},
	{path: "compass-runtime-agent", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "connection-token-handler", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "connectivity-certs-controller", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "connector-service", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release17),
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "console-backend-service", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release16),
			jobsuite.RunIfChanged("components/console-backend-service/main.go", "scripts/go-dep.mk"),
		},
	},
	{path: "dex-static-user-configurer", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "etcd-tls-setup-job", image: tester.ImageGolangBuildpack1_11},
	{path: "event-bus", image: tester.ImageGolangBuildpack1_11,
		additionalOptions: []jobsuite.Option{
			jobsuite.Until(releases.Release18),
		}},
	{path: "event-service", image: tester.ImageGolangBuildpack1_11,
		additionalOptions: []jobsuite.Option{
			jobsuite.Until(releases.Release18),
		}},
	{path: "iam-kubeconfig-service", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "istio-kyma-patch", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "k8s-dashboard-proxy", image: tester.ImageGolangBuildpack1_11},
	{path: "function-controller", image: tester.ImageGolangKubebuilderBuildpackLatest,
		additionalOptions: []jobsuite.Option{
			jobsuite.Until(releases.Release18),
		}},
	{path: "kubeless-images/nodejs", image: tester.ImageGolangBuildpack1_11,
		additionalOptions: []jobsuite.Option{
			jobsuite.Until(releases.Release18),
		}},
	{path: "kyma-operator", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "service-binding-usage-controller", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "xip-patch", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release17),
		},
	},
	{path: "health-check-proxy", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.Since(releases.Release19),
			jobsuite.Optional(),
		},
	},

	{path: "function-controller", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release19),
		},
	},
	{path: "event-service", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release19),
		},
	},
	{path: "event-bus", image: tester.ImageBootstrap20181204, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.Since(releases.Release19),
		},
	},
}

func TestComponentJobs(t *testing.T) {
	for _, component := range components {
		t.Run(component.path, func(t *testing.T) {
			opts := []jobsuite.Option{
				jobsuite.Component(component.path, component.image),
				jobsuite.KymaRepo(),
				jobsuite.AllReleases(),
			}
			opts = append(opts, component.additionalOptions...)
			cfg := jobsuite.NewConfig(opts...)
			suite := component.suite
			if suite == nil {
				suite = tester.NewComponentSuite
			}
			suite(cfg).Run(t)
		})
	}
}
