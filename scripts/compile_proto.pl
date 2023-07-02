#!/usr/bin/perl

my $go_module = "github.com/isd-sgcu/rpkm66-backend";
my $proto_package = "internal/proto";

my $source_dir = scalar(@ARGV) < 1 ? "*" : "${ARGV[0]}";
my $all_proto_files = `find ${source_dir} -name *.proto -type f`;

my $go_proto_package = "$go_module/$proto_package";

my @proto_files = split /\R/, $all_proto_files;
my @mapped = map { 
    (my $remove_prefix = $_) =~ s/\Q$source_dir\///;

    (my $go_package = $remove_prefix) =~ s/\/\w*\.proto//;

    ("--go_opt=M$remove_prefix=$go_proto_package/$go_package", "--go-grpc_opt=M$remove_prefix=$go_proto_package/$go_package");
} @proto_files;

my $proto_path = $source_dir == "*" ? "." : $source_dir;

my @cmd_prefix = (
    "protoc",
    "--go_out=.",
    "--go_opt=module=$go_module",
    "--go-grpc_out=.",
    "--go-grpc_opt=module=$go_module",
    "--proto_path=$proto_path",
    @mapped
);

my $cmd = (join " ", @cmd_prefix) . " " . (join " ", @proto_files);
# print $cmd;
system $cmd;