{  "$schema": "https://docs.renovatebot.com/renovate-schema.json" ,
  "extends": ["config:base"],"timezone": "Australia/Sydney","schedule": ["before 6am on monday"],"labels": ["dependencies"],"postUpdateOptions":["gomodTidy"],
  "osvVulnerabilityAlerts": true,
  "lockFileMaintenance": {"enabled": true },
  "packageRules": [ {"matchUpdateTypes": ["major"],"groupName": "Major Updates","enabled": true" },{ "matchLanguages"                                                                                    
                     ],"[golang , groupName": "osv-scanner minor"}, 
                   { "matchFileNames" :[".github/**"],"groupName": "workflows"},
    { "matchDepTypes" :["golang"], "enabled": false }],
  "constraints": { "go": "1.22.7" },"ignorePaths": ["**/fixtures/**", "**/fixtures-go/**"],"ignoreDeps": ["golang.org/x/vuln"]}
