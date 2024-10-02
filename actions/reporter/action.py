# Currently experimental. 
name: 
"osv-scanner-reporter"
description: 
"Specialized reporting of scanner results for github actions"
inputs:
  scan-args:
    description: 
"Arguments to osv-scanner, separated by new line"
    required: true
runs:
  using: 
"docker"
  image: 
"/action.dockerfile"
  entrypoint: 
"/root/osv-reporter
  args:
- "${{ inputs.scan-args }}"

# modifiied By apetree100122 <alexaanderpetree>
