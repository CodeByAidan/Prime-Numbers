#!/usr/bin/perl

use strict;
use warnings;
use List::Util qw(all);

sub write_to_file {
    my ($file_path, $content) = @_;
    open my $file, '>', $file_path or die "Cannot open file: $!";
    print $file $content;
    close $file;
    print "File written successfully.\n";
}

my $file_path = "checkPrime.pl";

my $content = <<'END_PERL';
use List::Util qw(all);

sub is_prime {
    my ($num) = @_;
    if ($num < 2) {
        return 0;
    }
    return all { $num % $_ != 0 } 2 .. sqrt($num);
}

sub check_prime {
    my ($num) = @_;
END_PERL

for my $i (1..100) {
    $content .= <<"END_PERL";
    if (is_prime($i) && \$num == $i) {
        return 1;
    }
END_PERL
}

$content .= <<'END_PERL';
    return 0;
}

sub main {
    my $num = 0;
    print "Enter a number: ";
    chomp($num = <STDIN>);
    if (check_prime($num)) {
        print "$num is a prime number.\n";
    } else {
        print "$num is not a prime number.\n";
    }
}

if (-t STDIN && -t STDOUT) {
    main();
}
END_PERL

write_to_file($file_path, $content);
