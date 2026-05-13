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
      # The actual service
      simple-service = pkgs.pkgsStatic.buildGoModule {
        name = "simple-service";
        version = "0.1.0";
        src = ./.;
        # There are only stdlib deps
        vendorHash = null;
      };
      # The corresponding docker image
      dockerImage = pkgs.dockerTools.buildLayeredImage {
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
    in
    {
      packages.${system} = {
        default = simple-service;
        docker-image = dockerImage;
      };
      
      apps.${system} = {
        push-docker-image = {
          type = "app";
          program = "${pkgs.writeShellApplication {
            name = "push-docker-image";
            runtimeInputs = [
              pkgs.skopeo
            ];
            text = ''
              GIT_TAG=$(git describe --tags --always 2>/dev/null)
              
              TAGGED_IMAGE="simple-service:$GIT_TAG"
              echo "Pushing docker image '$TAGGED_IMAGE' to registry '$REGISTRY'"

              skopeo copy docker-archive:${dockerImage} docker://"''${REGISTRY}"/simple-service:"$GIT_TAG" \
                --policy <(echo '{"default":[{"type":"insecureAcceptAnything"}]}') \
                --dest-tls-verify=false \
                --dest-creds="''${REGISTRY_USER}:''${REGISTRY_PASSWORD}"
            '';
          }}/bin/push-docker-image";
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
