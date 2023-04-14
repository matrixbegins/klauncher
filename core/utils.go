package core

import (
	"os"

	"github.com/joho/godotenv"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetContainerEnvVars() []apiv1.EnvVar {
	return []apiv1.EnvVar{
		{Name: "PROFILER_NAME", Value: GoDotEnvVariable("PROFILER_NAME")},
		{Name: "PROFILER_CRON_EXP", Value: GoDotEnvVariable("PROFILER_CRON_EXP")},
	}
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func GetJobSpec() *v1.Job {

	job := &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "counter-job",
		},
		Spec: v1.JobSpec{
			TTLSecondsAfterFinished: func(i int32) *int32 { return &i }(60),
			BackoffLimit:            func(i int32) *int32 { return &i }(1),
			Completions:             func(i int32) *int32 { return &i }(1),
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					OS: &apiv1.PodOS{Name: apiv1.Linux},
					Containers: []apiv1.Container{
						{
							Name: "counter-job", Image: GoDotEnvVariable("PROFILER_IMAGE"), Env: GetContainerEnvVars(),
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
		},
	}
	return job
}

func GetCronJobSpec() *v1.CronJob {

	job := &v1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name: "counter-cron-job",
		},
		Spec: v1.CronJobSpec{
			Schedule:               GoDotEnvVariable("PROFILER_CRON_EXP"),
			ConcurrencyPolicy:      v1.ForbidConcurrent,
			FailedJobsHistoryLimit: func(i int32) *int32 { return &i }(3),
			JobTemplate: v1.JobTemplateSpec{
				Spec: v1.JobSpec{
					TTLSecondsAfterFinished: func(i int32) *int32 { return &i }(60),
					BackoffLimit:            func(i int32) *int32 { return &i }(1),
					Completions:             func(i int32) *int32 { return &i }(1),
					Template: apiv1.PodTemplateSpec{
						Spec: apiv1.PodSpec{
							OS: &apiv1.PodOS{Name: apiv1.Linux},
							Containers: []apiv1.Container{
								{
									Name: "counter-cron", Image: GoDotEnvVariable("PROFILER_IMAGE"), Env: GetContainerEnvVars(),
								},
							},
							RestartPolicy: apiv1.RestartPolicyNever,
						},
					},
				},
			},
		},
	}

	return job
}
