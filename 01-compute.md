Demonstrate VM creation through CLI.

GCP VM (Tokyo region):

``` sh
# Create:
$ gcloud compute instances create zz1 --machine-type e2-micro --zone asia-northeast1-a

# Try connecting using browser SSH.

# Delete:
$ gcloud compute instances delete zz1 --zone asia-northeast1-a --quiet
```

AWS EC2 (Tokyo region):
``` sh
# Create template:
$ aws ec2 create-launch-template \
  --launch-template-name zz-lt \
  --version-description version1 \
  --launch-template-data '{"ImageId":"ami-0f75d1a8c9141bd00","InstanceType":"t2.micro"}'

# Create a VM using auto-scaling-group:
$ aws autoscaling create-auto-scaling-group \
  --auto-scaling-group-name zz-asg \
  --launch-template LaunchTemplateName=zz-lt,Version='1' \
  --min-size 1 \
  --max-size 1 \
  --availability-zones ap-northeast-1a

# 1) Connect using browser SSH.
# 2) Delete an instance to confirm autoscaling.

# Delete auto-scaling-group:
$ aws autoscaling delete-auto-scaling-group --auto-scaling-group-name zz-asg --force-delete

# Delete launch template:
$ aws ec2 delete-launch-template --launch-template-name zz-lt
```
