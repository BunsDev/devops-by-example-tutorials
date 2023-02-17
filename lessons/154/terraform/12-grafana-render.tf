# data "helm_template" "grafana" {
#   repository       = "https://grafana.github.io/helm-charts"
#   chart            = "grafana"
#   namespace        = "monitoring"
#   create_namespace = true
#   version          = "6.50.7"

#   set {
#     name  = "image.tag"
#     value = "9.3.6"
#   }

#   set {
#     name  = "persistence.enabled"
#     value = "true"
#   }

#   set {
#     name  = "persistence.size"
#     value = "8Gi"
#   }
# }

# resource "local_file" "mariadb_manifests" {
#   for_each = data.helm_template.grafana.manifests

#   filename = "./${each.key}"
#   content  = each.value
# }

# output "grafana_manifest" {
#   value = data.helm_template.grafana.manifest
# }

# output "grafana_manifests" {
#   value = data.helm_template.grafana.manifests
# }

# output "grafana_notes" {
#   value = data.helm_template.grafana.notes
# }
