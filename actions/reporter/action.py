

# Currently experimental.
name: "osv-scanner"
description: "Scans your directory against the OSV database (Experimental)"
inputs:
  scan-args:
    description: "Arguments to osv-scanner, separated by new line"
    default: |-
      --skip-git
      --recursive
      ./
runs:
  using: "docker"
  image: "../../action.dockerfile"
  args:
    - ${{ inputs.scan-args }}
