set -x
CSI_DRIVER_NAME=$1
DEPLOYEMENT_NAME=$(oc get pods -A -o go-template='{{ range .items}}{{ $alllabels := .metadata.labels}}{{ range .spec.containers }}{{ range .args }}{{if eq . "--driver-name='${CSI_DRIVER_NAME}'"}}{{ range $label,$value := $alllabels}}{{if eq $label "app.kubernetes.io/managed-by"}}{{$value}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}')
OPERATOR_NAME=$(oc get deployment  $DEPLOYEMENT_NAME  -o go-template='{{ range $label,$value := .metadata.labels}} {{$label}} {{print "\n"}} {{end}}' |grep "operators.coreos.com"| sed "s#operators.coreos.com/##g"|sed 's/ //g')
SUBSCRIPTION_NAME=$(oc get operator $OPERATOR_NAME -o go-template='{{ range .status.components.refs }} {{if eq .kind "Subscription"}} {{.name}} {{end}} {{end}}'|sed 's/ //g')
oc get subscription $SUBSCRIPTION_NAME -ogo-template={{.spec.source}} {{.spec.name}}'