{
  description = "Simple golang http to demo CD.";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.11";
  };

  outputs =
    { self, nixpkgs, ... }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
      simple-service = pkgs.pkgsStatic.buildGoModule {
        name = "simple-service";
        version = "0.1.0";
        src = ./.;
        # There are only stdlib deps
        vendorHash = null;
      };
    in
    {
      packages.${system} = {
        default = simple-service;
        docker-image = pkgs.dockerTools.buildLayeredImage {
          name = "simple-service";
          tag = "latest";
          
          contents = [
            simple-service
          ];
          
          config = {
            Cmd = [
              "${simple-service}/bin/simple-service"
            ];
          };
        };
      };

      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          delve
          go
          gotools
          gotestsum
          gopls
          #
          nixd
        ];
      };
    };
}
