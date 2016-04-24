# ghenga

[![Join the chat at https://gitter.im/ghenga/ghenga](https://badges.gitter.im/ghenga/ghenga.svg)](https://gitter.im/ghenga/ghenga?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
The ghenga micro CRM

## Getting Started

Install `vagrant` and execute `vagrant up` to download and provision the
development environment. Then use `vagrant ssh` to log into the VM.

The default vagrant base box is `ubuntu/xerus64`, if you need to use Parallels
instead of VirtualBox, you can provision and use a different base image as
follows:

```shell
export GHENGA_VAGRANT_BOX=boxcutter/ubuntu1604
vagrant up
```

While the VirtualBox download is running, you can clone the [ghenga-ui
repository](https://github.com/ghenga/ghenga-ui):

```shell
git clone https://github.com/ghenga/ghenga-ui
```

Afterwards, log into the virtual machine:

```
vagrant ssh
```

Then run the following commands to build and start the ghenga server:

```shell
cd ghenga
gb build
bin/ghenga serve --public ghenga-ui/app
```
