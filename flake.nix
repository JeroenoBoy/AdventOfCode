{
  description = "AOC perfomance measurer";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }: utils.lib.eachDefaultSystem (system: 
    let
      pkgs = (import nixpkgs) { inherit system; };
    in {
      devShell = pkgs.mkShell {
        name = "AOC performance shell";
        buildInputs = with pkgs; [ go gopls cmake nodejs_22 ];

        shellHook = ''
          echo ""
          go version
          echo "
          Welcome to the AOC perf shell 🚀!
          ";
        '';
      };
    }
  );
}
