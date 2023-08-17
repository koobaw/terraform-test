output "ip" {
    description = "IP address of the instance"
    value = google_compute_instance.server.network_interface.0.network_ip
}

output "instance_id" {
    description = "ID of GCE instance"
    value = google_compute_instance.server.id
}

