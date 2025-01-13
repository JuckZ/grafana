//
// This file is generated by grafana-app-sdk
// DO NOT EDIT
//

package apis

import (
	"encoding/json"

	"github.com/grafana/grafana-app-sdk/app"
)

var (
	rawSchemaInvestigationv1alpha1     = []byte(`{"spec":{"description":"spec is the schema of our resource. The spec should include all the user-ediable information for the kind.","properties":{"items":{"items":{"properties":{"dataQuery":{"description":"dataQuery contains the query used to generate this item.","oneOf":[{"allOf":[{"required":["refId","datasource","expr"]},{"not":{"anyOf":[{"required":["refId","datasource","expr"]}]}}]},{"allOf":[{"required":["refId","datasource","expr"]},{"not":{"anyOf":[{"required":["refId","datasource","expr"]}]}}]}],"properties":{"datasource":{"description":"datasource is the datasource of the query.","properties":{"apiVersion":{"type":"string"},"name":{"type":"string"},"type":{"type":"string"},"uid":{"type":"string"}},"required":["uid","type","apiVersion","name"],"type":"object"},"expr":{"description":"expr is the expression of the query.","type":"string"},"maxLines":{"description":"maxLines (optional) is used to limit the number of log rows returned.","format":"int64","type":"integer"},"refId":{"description":"refId is the reference ID of the query.","type":"string"}},"type":"object"},"iconPath":{"description":"iconPath (optional) is the path to the icon for the item.","type":"string"},"id":{"type":"string"},"note":{"description":"note (optional) is a comment on the item.","items":{"properties":{"authorUserID":{"type":"string"},"bodyMarkdown":{"type":"string"}},"required":["authorUserID","bodyMarkdown"],"type":"object"},"type":"array"},"origin":{"description":"origin is where the item was created from.","type":"string"},"queryType":{"description":"queryType is the type of the query used to generate this item.","enum":["logs","metrics"],"type":"string"},"timeRange":{"description":"timeRange (optional) is the time range of the item.","properties":{"from":{"type":"number"},"to":{"type":"number"}},"required":["from","to"],"type":"object"},"title":{"type":"string"},"type":{"description":"type is the type of the item \"timeseries\", \"heatmap\", \"log-table\" (not an enum to allow for future extensions).","type":"string"},"url":{"description":"url is the URL to the item.","type":"string"}},"required":["id","title","type","url","origin","timeRange","queryType","dataQuery"],"type":"object"},"type":"array"},"status":{"enum":["open","closed"],"type":"string"},"title":{"type":"string"}},"required":["title","status","items"],"type":"object"},"status":{"properties":{"additionalFields":{"description":"additionalFields is reserved for future use","type":"object","x-kubernetes-preserve-unknown-fields":true},"operatorStates":{"additionalProperties":{"properties":{"descriptiveState":{"description":"descriptiveState is an optional more descriptive state field which has no requirements on format","type":"string"},"details":{"description":"details contains any extra information that is operator-specific","type":"object","x-kubernetes-preserve-unknown-fields":true},"lastEvaluation":{"description":"lastEvaluation is the ResourceVersion last evaluated","type":"string"},"state":{"description":"state describes the state of the lastEvaluation.\nIt is limited to three possible states for machine evaluation.","enum":["success","in_progress","failed"],"type":"string"}},"required":["lastEvaluation","state"],"type":"object"},"description":"operatorStates is a map of operator ID to operator state evaluations.\nAny operator which consumes this kind SHOULD add its state evaluation information to this field.","type":"object"}},"type":"object","x-kubernetes-preserve-unknown-fields":true}}`)
	versionSchemaInvestigationv1alpha1 app.VersionSchema
	_                                  = json.Unmarshal(rawSchemaInvestigationv1alpha1, &versionSchemaInvestigationv1alpha1)
)

var appManifestData = app.ManifestData{
	AppName: "investigation",
	Group:   "investigation.grafana.app",
	Kinds: []app.ManifestKind{
		{
			Kind:       "Investigation",
			Scope:      "Namespaced",
			Conversion: false,
			Versions: []app.ManifestKindVersion{
				{
					Name:   "v1alpha1",
					Schema: &versionSchemaInvestigationv1alpha1,
				},
			},
		},
	},
}

func jsonToMap(j string) map[string]any {
	m := make(map[string]any)
	json.Unmarshal([]byte(j), &j)
	return m
}

func LocalManifest() app.Manifest {
	return app.NewEmbeddedManifest(appManifestData)
}

func RemoteManifest() app.Manifest {
	return app.NewAPIServerManifest("investigation")
}
