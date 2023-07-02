# sw-gittycat-server
GittyCat, as the name suggests, merges the principles of 'Git', the version control system, and 'Concatenate', reflecting the tool's ability to string together CI/CD pipelines. This open-source, Go-based, self-hosted CI/CD system effectively 'concatenates' your (not only) Git workflows, bringing together a secure and efficient, scalable CI/CD solution.

  ```
 /\_/\  
(=^.^=)
 > ^ <
  ```

# STILL UNDER DEVELOPMENT

An open-source, self-hosted CI/CD system that is designed to provide a seamless and secure experience for continuous integration and deployment. GittyCat enables the execution of automated workflows across multiple runners, making it an efficient and scalable solution for your CI/CD needs.

Developed in Go, GittyCat natively supports a wide range of architectures including 386, arm, arm64, and amd64. 
Its versatility extends to the operating systems it can function on, with robust support for FreeBSD, Linux, macOS, and Windows.

One of the key focus of GittyCat is its preference for SSH-based cloning. This not only enhances security but also enables it to work seamlessly with private repositories where other tools have complicated setup routines. 
It further expands its utility by supporting cloning from virtually any service that allows it, including GitHub, Gitea, and more.

GittyCat workflows and job descriptions are managed in a straightforward YAML format, providing users with an easy and understandable way to control their CI/CD processes.


## Current view of feature set and planned features:

| Feature | Not planned | Planned | Implemented |
| --- | --- | --- | --- |
| Clone HTTP |  |  | 🟢 |
| Clone HTTPS |  |  | 🟢 |
| Clone SSH |  |  | 🟢 |
| YAML workflows |  |  | 🟢 | (limited)
| Source stage |  |  | 🟢 |
| Build stage |  | 🟡 |  |
| Test stage |  | 🟡 |  |
| Deploy stage |  | 🟡 |  |
| Webhook trigger |  | 🟡 |  |
| Webinterface HTTPS/HTTPS |  | 🟡 |  |
| Gitea integration |  | 🟡 |  |
| Docker integration |  | 🟡 |  |
| OAuth  |  | 🟡 |  |


... More information soon!
