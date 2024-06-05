# This template is copy & pasted from https://nixos.wiki/wiki/Development_environment_with_nix-shell .
# nix is the best package manager. if you like docker then go write dockerconfig or whatever
{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    # nativeBuildInputs is usually what you want -- tools you need to run
    nativeBuildInputs = with pkgs.buildPackages; [ go ];
}

# to run nix, all you need to do is:
# 1. install nix
# 2. run `nix-shell` in this directory
