Name: hbm
Version: %{_version}
Release: %{_release}%{?dist}
Summary: Docker Engine Access Authorization Plugin
Group: Tools/Docker

License: GPL

URL: https://github.com/harbourmaster/hbm
Vendor: Kassisol
Packager: Kassisol <support@kassisol.com>

BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot

Source: hbm.tar.gz

%description
HBM is an application to authorize and manage authorized docker command.

%prep
%setup -n %{name}

%install
# install binary
install -d $RPM_BUILD_ROOT/%{_sbindir}
install -p -m 755 hbm $RPM_BUILD_ROOT/%{_sbindir}/

# add init scripts
install -d $RPM_BUILD_ROOT/%{_unitdir}
install -p -m 644 hbm.service $RPM_BUILD_ROOT/%{_unitdir}/hbm.service

# list files owned by the package here
%files
#%doc README.md
%{_sbindir}/hbm
/%{_unitdir}/hbm.service

%post
%systemd_post hbm

%preun
%systemd_preun hbm

%postun
rm -f %{_sbindir}/hbm

%systemd_postun_with_restart docker

%clean
rm -rf $RPM_BUILD_ROOT
