import pulumi
import pulumi_gcp as gcp

jammy = "projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20240208"
region = "us-west1"
zone = "us-west1-a"

with open("key.file", 'r', encoding='utf-8') as file:
  key = file.read()
with open("key.pub", 'r', encoding='utf-8') as file:
  pub = file.read()

with open("web-3000.service", 'r', encoding='utf-8') as file:
  web3000 = file.read()

startup_script = f"""#!/bin/bash

apt update

systemctl daemon-reload

cd /etc/systemd/system
echo -e "DATABASE_URL=todo" >> aa.conf
cat <<EOF > web-3000.service
{web3000}
EOF
useradd aa
mkdir /home/aa
chown aa:aa /home/aa
curl http
chmod +x /home/aa/web-3000
mkdir /Users/aa
chown -R aa:aa /home/aa
systemctl daemon-reload
systemctl enable web-3000
systemctl start web-3000
"""

static_ip = gcp.compute.Address("alb", region=region)

compute_instance = gcp.compute.Instance(
    "aa-aug-23-2024",
    machine_type="e2-micro-1",
    zone=zone,
    metadata_startup_script=startup_script,
    metadata={
      "enable-oslogin": "false",
      "ssh-keys": "root:ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICFVuvSZMJmYJcarbTq4rtKjdfwEPXI0T6bAa5k9/F6I root@devops",
    },
    boot_disk=gcp.compute.InstanceBootDiskArgs(
        initialize_params=gcp.compute.InstanceBootDiskInitializeParamsArgs(
            image=jammy,
            size=30,
            type="pd-ssd",
        )
    ),
    network_interfaces=[
        gcp.compute.InstanceNetworkInterfaceArgs(
            network="default",
            access_configs=[
                gcp.compute.InstanceNetworkInterfaceAccessConfigArgs(
                    nat_ip=static_ip.address,
                )
            ],
        )
    ],
    service_account=gcp.compute.InstanceServiceAccountArgs(
        scopes=["https://www.googleapis.com/auth/cloud-platform"],
    ),
    tags=["http-server", "https-server"]
)


pulumi.export("instanceName", compute_instance.name)
