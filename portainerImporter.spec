%define debug_package   %{nil}
%define _build_id_links none
%define _name portainerImporter
%define _prefix /opt
%define _version 1.00.00
%define _rel 0
%define _arch x86_64
%define _binaryname portainerImporter

Name:       portainerImporter
Version:    %{_version}
Release:    %{_rel}
Summary:    Import docker images to PortainerCE

Group:      Docker utils
License:    GPL2.0
URL:        

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
Import docker images to PortainerCE

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog