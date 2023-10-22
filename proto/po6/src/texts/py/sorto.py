#!/usr/bin/env python

"""
This python file is being written to alphabetize some entries 
"""
# data = [
#     {"1", "aix", "Interact LLM APIs"},
#     {"2", "alterx", "Subdomain wordlist gen using DSL."},
#     {"3", "asnmap", "Go Cli/Lib map network ranges ASN"},
#     {"4", "cdncheck", "ID tech with dns / ip."},
#     {"5", "chaos-client", "Go client comms Chaos dataset API."},
#     {"6", "cloudlist", "Multi-cloud Assets from Cloud"},
#     {"7", "dnsx", "DNS Toolkit"},
#     {"8", "httpx", "HTTP Toolkit"},
#     {"9", "interactsh-client", "Gathering Server"},
#     {"10", "interactsh-server", "Gathering Client"},
#     {"11", "katana", "Crawl/Spider Framework"},
#     {"12", "mapcidr", "Go Cli/Lib Ops Management CIDR"},
#     {"13", "naabu", "Go Por Scanner"},
#     {"14", "notify", "Output Streamer"},
#     {"15", "nuclei", "YAML/DSL Vuln Scanner"},
#     {"16", "pdtm", "Project Discovery Tool Manager"},
#     {"17", "proxify", "Swiss Army Knife Proxy"},
#     {"18", "shuffledns", "Go massdns wrapper"},
#     {"19", "simplehttpserver", "Go Alernative to python httpserver"},
#     {"20", "subfinder", "Subdomain Discovery"},
#     {"21", "tlsx", "Tls Grabber"},
#     {"22", "uncover", "Exposed host disvoery via search"},
# ]
data = [
    ["1", "aix", "Interact LLM APIs"],
    ["2", "alterx", "Subdomain wordlist gen using DSL."],
    ["3", "asnmap", "Go Cli/Lib map network ranges ASN"],
    ["4", "cdncheck", "ID tech with dns / ip."],
    ["5", "chaos-client", "Go client comms Chaos dataset API."],
    ["6", "cloudlist", "Multi-cloud Assets from Cloud"],
    ["7", "dnsx", "DNS Toolkit"],
    ["8", "httpx", "HTTP Toolkit"],
    ["9", "interactsh-client", "Gathering Server"],
    ["10", "interactsh-server", "Gathering Client"],
    ["11", "katana", "Crawl/Spider Framework"],
    ["12", "mapcidr", "Go Cli/Lib Ops Management CIDR"],
    ["13", "naabu", "Go Por Scanner"],
    ["14", "notify", "Output Streamer"],
    ["15", "nuclei", "YAML/DSL Vuln Scanner"],
    ["16", "pdtm", "Project Discovery Tool Manager"],
    ["17", "proxify", "Swiss Army Knife Proxy"],
    ["18", "shuffledns", "Go massdns wrapper"],
    ["19", "simplehttpserver", "Go Alernative to python httpserver"],
    ["20", "subfinder", "Subdomain Discovery"],
    ["21", "tlsx", "Tls Grabber"],
    ["22", "uncover", "Exposed host disvoery via search"],
]

sorted_data = sorted(data, key=lambda x: x[1])

for item in sorted_data:
    print(item)
