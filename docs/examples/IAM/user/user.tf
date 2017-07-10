variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
	tenancy_ocid = "${var.tenancy_ocid}"
	user_ocid = "${var.user_ocid}"
	fingerprint = "${var.fingerprint}"
	private_key_path = "${var.private_key_path}"
}

resource "baremetal_identity_user" "tf_user" {
	name = "tf_user"
	description = "A user I'm managing with Terraform"
}

resource "baremetal_identity_ui_password" "tf_password" {
	user_id = "${baremetal_identity_user.tf_user.id}"
}

output "UserUIPassword" {
	sensitive = false
	value = ["${baremetal_identity_ui_password.tf_password.password}"]
}
