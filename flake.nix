{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { nixpkgs, flake-utils, ... }: flake-utils.lib.eachDefaultSystem (system:
    let pkgs = nixpkgs.legacyPackages.${system}; in {

      devShell = pkgs.mkShellNoCC {
        name = "JourniCal devshell";
        buildInputs = with pkgs; [
          go
          nodejs_22
          gnumake
        ];
        shellHook = '''';
      };
    });
}
