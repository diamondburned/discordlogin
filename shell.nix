{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation rec {
	name = "gtkcord3";

	buildInputs = with pkgs; [
		gnome3.gtk gnome3.webkitgtk
		pkgconfig go
	];
}
