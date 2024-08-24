import pulumi
import pulumi_gcp as gcp

jammy = "projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20240208"
region = "us-west1"
zone = "us-west1-a"

with open("web-3000.service", 'r', encoding='utf-8') as file:
  web3000 = file.read()

startup_script = f"""#!/bin/bash
cd /etc/systemd/system
echo -e "DATABASE_URL=todo" >> aa.conf
cat <<EOF > web-3000.service
{web3000}
EOF
useradd aa
mkdir /home/aa
chown aa:aa /home/aa
cd /home/aa
curl "https://objects.githubusercontent.com/github-production-release-asset-2e65be/846781077/9df9114e-af19-4b7e-bec0-70c4599dd5bb?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=releaseassetproduction%2F20240824%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240824T011444Z&X-Amz-Expires=300&X-Amz-Signature=f3a65e4862c48080b6ea5ff92ec7df0a3bea006c664011333517a41715df0225&X-Amz-SignedHeaders=host&actor_id=0&key_id=0&repo_id=846781077&response-content-disposition=attachment%3B%20filename%3Dinspiredby2&response-content-type=application%2Foctet-stream" -o web-3000
chmod +x /home/aa/web-3000
mkdir /Users/aa
chown -R aa:aa /home/aa
systemctl daemon-reload
systemctl enable web-3000
systemctl start web-3000
"""

static_ip = gcp.compute.Address("alb2", region=region)

compute_instance = gcp.compute.Instance(
    "aa-aug-23-2024",
    machine_type="e2-micro",
    zone=zone,
    metadata_startup_script=startup_script,
    metadata={
      "enable-oslogin": "false",
      "ssh-keys": "root:ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIH1W5dxq6/7D3p+dIfajSv9RWcEKm5OpCu7rOm9rGEIx root@aa.local",
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
