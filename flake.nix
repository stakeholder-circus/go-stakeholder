{
  description = "stakeholder-circus go-stakeholder";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-darwin" "x86_64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in {
      devShells = forAllSystems (system:
        let pkgs = import nixpkgs { inherit system; };
        in {
          default = pkgs.mkShell {
            packages = with pkgs; [ go golangci-lint docker ];
          };
        });
      apps = forAllSystems (system:
        let pkgs = import nixpkgs { inherit system; };
            mk = name: text: {
              type = "app";
              program = "${pkgs.writeShellScript name text}";
            };
        in {
          build = mk "build" ''go build ./...'';
          test = mk "test" ''go test ./...'';
          check = mk "check" ''files=$(find . -name '*.go' -not -path './.git/*' -print); if [ -n "$files" ] && [ -n "$(gofmt -l $files)" ]; then gofmt -l $files; exit 1; fi; go vet ./...; go build ./...; go test ./...; golangci-lint run'';
          format = mk "format" ''files=$(find . -name '*.go' -not -path './.git/*' -print); if [ -n "$files" ] && [ -n "$(gofmt -l $files)" ]; then gofmt -l $files; exit 1; fi'';
        });
    };
}
