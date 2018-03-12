Name: hbm
Version: %{_version}
Release: %{_release}%{?dist}
Summary: Docker Engine Access Authorization Plugin
Group: Tools/Docker

License: GPLv3+

URL: https://github.com/kassisol/hbm
Vendor: Kassisol
Packager: Kassisol <support@kassisol.com>

BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot

Source: hbm.tar.gz

%description
HBM is an authorization plugin for docker commands.

%prep
%setup -n %{name}

%install
# install binary
install -d $RPM_BUILD_ROOT/%{_sbindir}
install -p -m 755 hbm $RPM_BUILD_ROOT/%{_sbindir}/

# add init scripts
install -d $RPM_BUILD_ROOT/%{_unitdir}
install -p -m 644 hbm.service $RPM_BUILD_ROOT/%{_unitdir}/hbm.service
install -p -m 644 hbm.socket $RPM_BUILD_ROOT/%{_unitdir}/hbm.socket

# add bash completions
install -d $RPM_BUILD_ROOT/usr/share/bash-completion/completions
install -p -m 644 shellcompletion/bash $RPM_BUILD_ROOT/usr/share/bash-completion/completions/hbm

# install manpages
install -d $RPM_BUILD_ROOT/%{_mandir}/man8
install -p -m 644 man/man8/*.8 $RPM_BUILD_ROOT/%{_mandir}/man8/

%files
#%doc README.md LICENSE
%{_sbindir}/hbm
/%{_unitdir}/hbm.service
/%{_unitdir}/hbm.socket
/usr/share/bash-completion/completions/hbm
/%{_mandir}/man8/*

%post
%systemd_post hbm.service
%systemd_post hbm.socket

%preun
%systemd_preun hbm.service
%systemd_preun hbm.socket

%clean
rm -rf $RPM_BUILD_ROOT
