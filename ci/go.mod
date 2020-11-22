module hello-app/ci

go 1.14

//replace github.com/tahirali-csc/task-executor-ci => "/Users/tahir/workspace/rnd-projects/product launch/task-executor-ci"

//GitHub Issue : https://github.com/kubernetes/client-go/issues/874
replace k8s.io/client-go => k8s.io/client-go v0.19.2

require github.com/tahirali-csc/task-executor-ci v0.0.2
