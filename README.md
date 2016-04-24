# ghenga
The ghenga micro CRM

## Getting Started

Install `vagrant` and execute `vagrant up` to download and provision the
development environment. Then use `vagrant ssh` to log into the VM.

The default vagrant base box is `ubuntu/xerus64`, if you need to use Parallels
instead of VirtualBox, you can provision and use a different base image as
follows:

    export GHENGA_VAGRANT_BOX=ffuenf/ubuntu-16.04-server-amd64
    vagrant up

ghenga can then be built and started with the following commands:

    cd ghenga
    gb build
    bin/ghenga serve
