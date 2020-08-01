package client

var templates = map[string]string{
	"action":        "e3sgR2VuZXJhdGVkV2FybmluZyB9fQoKcGFja2FnZSB7eyAuUGFja2FnZSB9fQoKaW1wb3J0ICgKCSJjb250ZXh0IgoJImVuY29kaW5nL2pzb24iCgkiZm10IgoJInN5bmMiCgkiaW8iCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3RvY29sIgoJcnBjY2xpZW50ICJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvcHJvdmlkZXJzL2FnZW50L21jb3JwYy9jbGllbnQiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3ZpZGVycy9hZ2VudC9tY29ycGMvcmVwbHlmbXQiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3ZpZGVycy9hZ2VudC9tY29ycGMvZGRsL2FnZW50IgopCgovLyB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RlciBwZXJmb3JtcyBhIFJQQyByZXF1ZXN0IHRvIHt7IC5BZ2VudE5hbWUgfCBUb0xvd2VyIH19I3t7IC5BY3Rpb25OYW1lIHwgVG9Mb3dlciB9fQp0eXBlIHt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdGVyIHN0cnVjdCB7CglyICAgICpyZXF1ZXN0ZXIKCW91dGMgY2hhbiAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQKfQoKLy8ge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQgaXMgdGhlIG91dHB1dCBmcm9tIHRoZSB7eyAuQWN0aW9uTmFtZSB8IFRvTG93ZXIgfX0gYWN0aW9uCnR5cGUge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQgc3RydWN0IHsKCWRldGFpbHMgKlJlc3VsdERldGFpbHMKCXJlcGx5ICAgbWFwW3N0cmluZ11pbnRlcmZhY2V7fQp9CgovLyB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdCBpcyB0aGUgcmVzdWx0IGZyb20gYSB7eyAuQWN0aW9uTmFtZSB8IFRvTG93ZXIgfX0gYWN0aW9uCnR5cGUge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQgc3RydWN0IHsKCWRkbCAqYWdlbnQuRERMCglzdGF0cyAgICpycGNjbGllbnQuU3RhdHMKCW91dHB1dHMgW10qe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQKCXJwY3JlcGxpZXMgW10qcmVwbHlmbXQuUlBDUmVwbHkKCW11IHN5bmMuTXV0ZXgKfQoKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQpIFJlbmRlclJlc3VsdHModyBpby5Xcml0ZXIsIGZvcm1hdCBSZW5kZXJGb3JtYXQsIGRpc3BsYXlNb2RlIERpc3BsYXlNb2RlLCB2ZXJib3NlIGJvb2wsIHNpbGVudCBib29sLCBsb2cgTG9nKSBlcnJvciB7CglkLm11LkxvY2soKQoJZGVmZXIgZC5tdS5VbmxvY2soKQoKCWlmIGQuc3RhdHMgPT0gbmlsIHsKCQlyZXR1cm4gZm10LkVycm9yZigicmVzdWx0IHN0YXRzIGlzIG5vdCBzZXQsIHJlc3VsdCB3YXMgbm90IGNvbXBsZXRlZCIpCgl9CgoJcmVzdWx0cyA6PSAmcmVwbHlmbXQuUlBDUmVzdWx0c3sKCQlBZ2VudDogICBkLnN0YXRzLkFnZW50KCksCgkJQWN0aW9uOiAgZC5zdGF0cy5BY3Rpb24oKSwKCQlSZXBsaWVzOiBkLnJwY3JlcGxpZXMsCgkJU3RhdHM6ICAgZC5zdGF0cywKCX0KCglhZGRsLCBlcnIgOj0gZC5kZGwuQWN0aW9uSW50ZXJmYWNlKGQuc3RhdHMuQWN0aW9uKCkpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gZXJyCgl9CgoJc3dpdGNoIGZvcm1hdCB7CgljYXNlIEpTT05Gb3JtYXQ6CgkJcmV0dXJuIHJlc3VsdHMuUmVuZGVySlNPTih3LCBhZGRsKQoJY2FzZSBUYWJsZUZvcm1hdDoKCQlyZXR1cm4gcmVzdWx0cy5SZW5kZXJUYWJsZSh3LCBhZGRsKQoJY2FzZSBUWFRGb290ZXI6CgkJcmVzdWx0cy5SZW5kZXJUWFRGb290ZXIodywgdmVyYm9zZSkKCQlyZXR1cm4gbmlsCglkZWZhdWx0OgoJCXJldHVybiByZXN1bHRzLlJlbmRlclRYVCh3LCBhZGRsLCB2ZXJib3NlLCBzaWxlbnQsIHJlcGx5Zm10LkRpc3BsYXlNb2RlKGRpc3BsYXlNb2RlKSwgbG9nKQoJfQp9CgovLyBTdGF0cyBpcyB0aGUgcnBjIHJlcXVlc3Qgc3RhdHMKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQpIFN0YXRzKCkgU3RhdHMgewoJcmV0dXJuIGQuc3RhdHMKfQoKLy8gUmVzdWx0RGV0YWlscyBpcyB0aGUgZGV0YWlscyBhYm91dCB0aGUgcmVxdWVzdApmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCkgUmVzdWx0RGV0YWlscygpICpSZXN1bHREZXRhaWxzIHsKCXJldHVybiBkLmRldGFpbHMKfQoKLy8gSGFzaE1hcCBpcyB0aGUgcmF3IG91dHB1dCBkYXRhCmZ1bmMgKGQgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSBIYXNoTWFwKCkgbWFwW3N0cmluZ11pbnRlcmZhY2V7fSB7CglyZXR1cm4gZC5yZXBseQp9CgovLyBKU09OIGlzIHRoZSBKU09OIHJlcHJlc2VudGF0aW9uIG9mIHRoZSBvdXRwdXQgZGF0YQpmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCkgSlNPTigpIChbXWJ5dGUsIGVycm9yKSB7CglyZXR1cm4ganNvbi5NYXJzaGFsKGQucmVwbHkpCn0KCi8vIFBhcnNlT3V0cHV0IHBhcnNlcyB0aGUgcmVzdWx0IHZhbHVlIGZyb20gdGhlIHt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19IGFjdGlvbiBpbnRvIHRhcmdldApmdW5jIChkICp7eyAkLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQpIFBhcnNle3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQodGFyZ2V0IGludGVyZmFjZXt9KSBlcnJvciB7CglqLCBlcnIgOj0gZC5KU09OKCkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBmbXQuRXJyb3JmKCJjb3VsZCBub3QgYWNjZXNzIHBheWxvYWQ6ICVzIiwgZXJyKQoJfQoKCWVyciA9IGpzb24uVW5tYXJzaGFsKGosIHRhcmdldCkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBmbXQuRXJyb3JmKCJjb3VsZCBub3QgdW5tYXJzaGFsIEpTT04gcGF5bG9hZDogJXMiLCBlcnIpCgl9CgoJcmV0dXJuIG5pbAp9CgovLyBEbyBwZXJmb3JtcyB0aGUgcmVxdWVzdApmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RlcikgRG8oY3R4IGNvbnRleHQuQ29udGV4dCkgKCp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdCwgZXJyb3IpIHsKCWRyZXMgOj0gJnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVzdWx0e2RkbDogZC5yLmNsaWVudC5kZGx9CgoJaGFuZGxlciA6PSBmdW5jKHByIHByb3RvY29sLlJlcGx5LCByICpycGNjbGllbnQuUlBDUmVwbHkpIHsKCQlvdXRwdXQgOj0gJnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0ewoJCQlyZXBseTogbWFrZShtYXBbc3RyaW5nXWludGVyZmFjZXt9KSwKCQkJZGV0YWlsczogJlJlc3VsdERldGFpbHN7CgkJCQlzZW5kZXI6ICBwci5TZW5kZXJJRCgpLAoJCQkJY29kZTogICAgaW50KHIuU3RhdHVzY29kZSksCgkJCQltZXNzYWdlOiByLlN0YXR1c21zZywKCQkJCXRzOiAgICAgIHByLlRpbWUoKSwKCQkJfSwKCQl9CgoJCWVyciA6PSBqc29uLlVubWFyc2hhbChyLkRhdGEsICZvdXRwdXQucmVwbHkpCgkJaWYgZXJyICE9IG5pbCB7CgkJCWQuci5jbGllbnQuZXJyb3JmKCJDb3VsZCBub3QgZGVjb2RlIHJlcGx5IGZyb20gJXM6ICVzIiwgcHIuU2VuZGVySUQoKSwgZXJyKQoJCX0KCgkJLy8gY2FsbGVyIHdhbnRzIGEgY2hhbm5lbCBzbyB3ZSBkb250IHJldHVybiBhIHJlc3Vsc2V0IHRvbyAobG90cyBvZiBtZW1vcnkgZXRjKQoJCS8vIHRoaXMgaXMgdW51c2VkIG5vdywgbm8gc3VwcG9ydCBmb3Igc2V0dGluZyBhIGNoYW5uZWwKCQlpZiBkLm91dGMgIT0gbmlsIHsKCQkJZC5vdXRjIDwtIG91dHB1dAoJCQlyZXR1cm4KCQl9CgoJCS8vIGVsc2UgcHJlcGFyZSBvdXIgcmVzdWx0IHNldAoJCWRyZXMubXUuTG9jaygpCgkJZHJlcy5vdXRwdXRzID0gYXBwZW5kKGRyZXMub3V0cHV0cywgb3V0cHV0KQoJCWRyZXMucnBjcmVwbGllcyA9IGFwcGVuZChkcmVzLnJwY3JlcGxpZXMsICZyZXBseWZtdC5SUENSZXBseXsKCQkJU2VuZGVyOiAgIHByLlNlbmRlcklEKCksCgkJCVJQQ1JlcGx5OiByLAoJCX0pCgkJZHJlcy5tdS5VbmxvY2soKQoJfQoKCXJlcywgZXJyIDo9IGQuci5kbyhjdHgsIGhhbmRsZXIpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBlcnIKCX0KCglkcmVzLnN0YXRzID0gcmVzCgoJcmV0dXJuIGRyZXMsIG5pbAp9CgovLyBFYWNoT3V0cHV0IGl0ZXJhdGVzIG92ZXIgYWxsIHJlc3VsdHMgcmVjZWl2ZWQKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQpIEVhY2hPdXRwdXQoaCBmdW5jKHIgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSkgewoJZm9yIF8sIHJlc3AgOj0gcmFuZ2UgZC5vdXRwdXRzIHsKCQloKHJlc3ApCgl9Cn0KCnt7IHJhbmdlICRuYW1lLCAkaW5wdXQgOj0gLk9wdGlvbmFsSW5wdXRzIH19Ci8vIHt7ICRuYW1lIHwgU25ha2VUb0NhbWVsIH19IGlzIGFuIG9wdGlvbmFsIGlucHV0IHRvIHRoZSB7eyAkLkFjdGlvbk5hbWUgfCBUb0xvd2VyIH19IGFjdGlvbgovLwovLyBEZXNjcmlwdGlvbjoge3sgJGlucHV0LkRlc2NyaXB0aW9uIH19CmZ1bmMgKGQgKnt7ICQuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3Rlcikge3sgJG5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0odiB7eyBDaG9yaWFUeXBlVG9Hb1R5cGUgJGlucHV0LlR5cGUgfX0pICp7eyAkLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXF1ZXN0ZXIgewoJZC5yLmFyZ3NbInt7ICRuYW1lIHwgVG9Mb3dlciB9fSJdID0gdgoKCXJldHVybiBkCn0Ke3sgZW5kIH19Cnt7IHJhbmdlICRuYW1lLCAkb3V0cHV0IDo9IC5PdXRwdXRzIH19Ci8vIHt7ICRuYW1lIHwgU25ha2VUb0NhbWVsIH19IGlzIHRoZSB2YWx1ZSBvZiB0aGUge3sgJG5hbWUgfX0gb3V0cHV0Ci8vCi8vIERlc2NyaXB0aW9uOiB7eyAkb3V0cHV0LkRlc2NyaXB0aW9uIH19CmZ1bmMgKGQgKnt7ICQuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCkge3sgJG5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0oKSB7eyBDaG9yaWFUeXBlVG9Hb1R5cGUgJG91dHB1dC5UeXBlIH19IHsKCXZhbCA6PSBkLnJlcGx5WyJ7eyAkbmFtZSB9fSJdCglyZXR1cm4ge3sgQ2hvcmlhVHlwZVRvVmFsT2ZUeXBlICRvdXRwdXQuVHlwZSB9fQp9Cnt7IGVuZCB9fQo=",
	"client":        "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImVuY29kaW5nL2Jhc2U2NCIKCSJlbmNvZGluZy9qc29uIgoJImZtdCIKCSJzeW5jIgoJInRpbWUiCgoJImNvbnRleHQiCgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9jaG9yaWEiCgljb3JlY2xpZW50ICJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvY2xpZW50L2NsaWVudCIKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvY29uZmlnIgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9zcnZjYWNoZSIKCXJwY2NsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3ZpZGVycy9hZ2VudC9tY29ycGMvY2xpZW50IgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjL2RkbC9hZ2VudCIKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvcHJvdG9jb2wiCgkiZ2l0aHViLmNvbS9zaXJ1cHNlbi9sb2dydXMiCikKCi8vIFN0YXRzIGFyZSB0aGUgc3RhdGlzdGljcyBmb3IgYSByZXF1ZXN0CnR5cGUgU3RhdHMgaW50ZXJmYWNlIHsKCUFnZW50KCkgc3RyaW5nCglBY3Rpb24oKSBzdHJpbmcKCUFsbCgpIGJvb2wKCU5vUmVzcG9uc2VGcm9tKCkgW11zdHJpbmcKCVVuZXhwZWN0ZWRSZXNwb25zZUZyb20oKSBbXXN0cmluZwoJRGlzY292ZXJlZENvdW50KCkgaW50CglEaXNjb3ZlcmVkTm9kZXMoKSAqW11zdHJpbmcKCUZhaWxDb3VudCgpIGludAoJT0tDb3VudCgpIGludAoJUmVzcG9uc2VzQ291bnQoKSBpbnQKCVB1Ymxpc2hEdXJhdGlvbigpICh0aW1lLkR1cmF0aW9uLCBlcnJvcikKCVJlcXVlc3REdXJhdGlvbigpICh0aW1lLkR1cmF0aW9uLCBlcnJvcikKCURpc2NvdmVyeUR1cmF0aW9uKCkgKHRpbWUuRHVyYXRpb24sIGVycm9yKQp9CgovLyBOb2RlU291cmNlIGRpc2NvdmVycyBub2Rlcwp0eXBlIE5vZGVTb3VyY2UgaW50ZXJmYWNlIHsKCVJlc2V0KCkKCURpc2NvdmVyKGN0eCBjb250ZXh0LkNvbnRleHQsIGZ3IENob3JpYUZyYW1ld29yaywgZmlsdGVycyBbXUZpbHRlckZ1bmMpIChbXXN0cmluZywgZXJyb3IpCn0KCi8vIENob3JpYUZyYW1ld29yayBpcyB0aGUgQ2hvcmlhIGZyYW1ld29yawp0eXBlIENob3JpYUZyYW1ld29yayBpbnRlcmZhY2UgewoJTG9nZ2VyKHN0cmluZykgKmxvZ3J1cy5FbnRyeQoJU2V0TG9nZ2VyKCpsb2dydXMuTG9nZ2VyKQoJQ29uZmlndXJhdGlvbigpICpjb25maWcuQ29uZmlnCglOZXdNZXNzYWdlKHBheWxvYWQgc3RyaW5nLCBhZ2VudCBzdHJpbmcsIGNvbGxlY3RpdmUgc3RyaW5nLCBtc2dUeXBlIHN0cmluZywgcmVxdWVzdCAqY2hvcmlhLk1lc3NhZ2UpIChtc2cgKmNob3JpYS5NZXNzYWdlLCBlcnIgZXJyb3IpCglOZXdSZXBseUZyb21UcmFuc3BvcnRKU09OKHBheWxvYWQgW11ieXRlLCBza2lwdmFsaWRhdGUgYm9vbCkgKG1zZyBwcm90b2NvbC5SZXBseSwgZXJyIGVycm9yKQoJTmV3VHJhbnNwb3J0RnJvbUpTT04oZGF0YSBzdHJpbmcpIChtZXNzYWdlIHByb3RvY29sLlRyYW5zcG9ydE1lc3NhZ2UsIGVyciBlcnJvcikKCU1pZGRsZXdhcmVTZXJ2ZXJzKCkgKHNlcnZlcnMgc3J2Y2FjaGUuU2VydmVycywgZXJyIGVycm9yKQoJTmV3Q29ubmVjdG9yKGN0eCBjb250ZXh0LkNvbnRleHQsIHNlcnZlcnMgZnVuYygpIChzcnZjYWNoZS5TZXJ2ZXJzLCBlcnJvciksIG5hbWUgc3RyaW5nLCBsb2dnZXIgKmxvZ3J1cy5FbnRyeSkgKGNvbm4gY2hvcmlhLkNvbm5lY3RvciwgZXJyIGVycm9yKQoJTmV3UmVxdWVzdElEKCkgKHN0cmluZywgZXJyb3IpCglDZXJ0bmFtZSgpIHN0cmluZwp9CgovLyBGaWx0ZXJGdW5jIGNhbiBnZW5lcmF0ZSBhIENob3JpYSBmaWx0ZXIKdHlwZSBGaWx0ZXJGdW5jIGZ1bmMoZiAqcHJvdG9jb2wuRmlsdGVyKSBlcnJvcgoKLy8gUmVuZGVyRm9ybWF0IGlzIHRoZSBmb3JtYXQgdXNlZCBieSB0aGUgUmVuZGVyUmVzdWx0cyBoZWxwZXIKdHlwZSBSZW5kZXJGb3JtYXQgaW50Cgpjb25zdCAoCgkvLyBKU09ORm9ybWF0IHJlbmRlcnMgdGhlIHJlc3VsdHMgYXMgYSBKU09OIGRvY3VtZW50CglKU09ORm9ybWF0IFJlbmRlckZvcm1hdCA9IGlvdGEKCgkvLyBUZXh0Rm9ybWF0IHJlbmRlcnMgdGhlIHJlc3VsdHMgYXMgYSBDaG9yaWEgdHlwaWNhbCByZXN1bHQgc2V0IGluIGxpbmUgd2l0aCBjaG9yaWEgcmVxIG91dHB1dAoJVGV4dEZvcm1hdAoKCS8vIFRhYmxlRm9ybWF0IHJlbmRlcnMgYWxsIHN1Y2Nlc3NmdWwgcmVzcG9uc2VzIGluIGEgdGFibGUKCVRhYmxlRm9ybWF0CgoJLy8gVFhURm9vdGVyIHJlbmRlcnMgb25seSB0aGUgcmVxdWVzdCBzdW1tYXJ5IHN0YXRpc3RpY3MKCVRYVEZvb3RlcgopCgovLyBEaXNwbGF5TW9kZSBvdmVycmlkZXMgdGhlIERETCBkaXNwbGF5IGhpbnRzCnR5cGUgRGlzcGxheU1vZGUgdWludDgKCmNvbnN0ICgKCS8vIERpc3BsYXlEREwgc2hvd3MgcmVzdWx0cyBiYXNlZCBvbiB0aGUgY29uZmlndXJhdGlvbiBpbiB0aGUgRERMIGZpbGUKCURpc3BsYXlEREwgPSBEaXNwbGF5TW9kZShpb3RhKQoJLy8gRGlzcGxheU9LIHNob3dzIG9ubHkgcGFzc2luZyByZXN1bHRzCglEaXNwbGF5T0sKCS8vIERpc3BsYXlGYWlsZWQgc2hvd3Mgb25seSBmYWlsZWQgcmVzdWx0cwoJRGlzcGxheUZhaWxlZAoJLy8gRGlzcGxheUFsbCBzaG93cyBhbGwgcmVzdWx0cwoJRGlzcGxheUFsbAoJLy8gRGlzcGxheU5vbmUgc2hvd3Mgbm8gcmVzdWx0cwoJRGlzcGxheU5vbmUKKQoKdHlwZSBMb2cgaW50ZXJmYWNlIHsKCURlYnVnZihmb3JtYXQgc3RyaW5nLCBhcmdzIC4uLmludGVyZmFjZXt9KQoJSW5mb2YoZm9ybWF0IHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkKCVdhcm5mKGZvcm1hdCBzdHJpbmcsIGFyZ3MgLi4uaW50ZXJmYWNle30pCglFcnJvcmYoZm9ybWF0IHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkKCUZhdGFsZihmb3JtYXQgc3RyaW5nLCBhcmdzIC4uLmludGVyZmFjZXt9KQoJUGFuaWNmKGZvcm1hdCBzdHJpbmcsIGFyZ3MgLi4uaW50ZXJmYWNle30pCn0KCi8vIHt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB0byB0aGUge3sgLkRETC5NZXRhZGF0YS5OYW1lIH19IGFnZW50CnR5cGUge3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHN0cnVjdCB7CglmdyAgICAgICAgICAgQ2hvcmlhRnJhbWV3b3JrCgljZmcgICAgICAgICAgICpjb25maWcuQ29uZmlnCglkZGwgICAgICAgICAgICphZ2VudC5EREwKCW5zICAgICAgICAgICAgTm9kZVNvdXJjZQoJY2xpZW50T3B0cyAgICAqaW5pdE9wdGlvbnMKCWNsaWVudFJQQ09wdHMgW11ycGNjbGllbnQuUmVxdWVzdE9wdGlvbgoJZmlsdGVycyAgICAgICBbXUZpbHRlckZ1bmMKCXRhcmdldHMgICAgICAgW11zdHJpbmcKCglzeW5jLk11dGV4Cn0KCi8vIE1ldGFkYXRhIGlzIHRoZSBhZ2VudCBtZXRhZGF0YQp0eXBlIE1ldGFkYXRhIHN0cnVjdCB7CglMaWNlbnNlICAgICBzdHJpbmcgYGpzb246ImxpY2Vuc2UiYAoJQXV0aG9yICAgICAgc3RyaW5nIGBqc29uOiJhdXRob3IiYAoJVGltZW91dCAgICAgaW50ICAgIGBqc29uOiJ0aW1lb3V0ImAKCU5hbWUgICAgICAgIHN0cmluZyBganNvbjoibmFtZSJgCglWZXJzaW9uICAgICBzdHJpbmcgYGpzb246InZlcnNpb24iYAoJVVJMICAgICAgICAgc3RyaW5nIGBqc29uOiJ1cmwiYAoJRGVzY3JpcHRpb24gc3RyaW5nIGBqc29uOiJkZXNjcmlwdGlvbiJgCn0KCi8vIE11c3QgY3JlYXRlIGEgbmV3IGNsaWVudCBhbmQgcGFuaWNzIG9uIGVycm9yCmZ1bmMgTXVzdChvcHRzIC4uLkluaXRpYWxpemF0aW9uT3B0aW9uKSAoY2xpZW50ICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIHsKCWMsIGVyciA6PSBOZXcob3B0cy4uLikKCWlmIGVyciAhPSBuaWwgewoJCXBhbmljKGVycikKCX0KCglyZXR1cm4gYwp9CgovLyBOZXcgY3JlYXRlcyBhIG5ldyBjbGllbnQgdG8gdGhlIHt7IC5EREwuTWV0YWRhdGEuTmFtZSB9fSBhZ2VudApmdW5jIE5ldyhvcHRzIC4uLkluaXRpYWxpemF0aW9uT3B0aW9uKSAoY2xpZW50ICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQsIGVyciBlcnJvcikgewoJYyA6PSAme3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50ewoJCWRkbDogICAgICAgICAgICZhZ2VudC5EREx7fSwKCQljbGllbnRSUENPcHRzOiBbXXJwY2NsaWVudC5SZXF1ZXN0T3B0aW9ue30sCgkJZmlsdGVyczogICAgICAgW11GaWx0ZXJGdW5jewoJCSAgICBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuQWdlbnRGaWx0ZXIoInt7IC5EREwuTWV0YWRhdGEuTmFtZSB9fSIpKSwKCQl9LAoJCWNsaWVudE9wdHM6ICZpbml0T3B0aW9uc3sKCQkJY2ZnRmlsZTogY2hvcmlhLlVzZXJDb25maWcoKSwKCQl9LAoJCXRhcmdldHM6IFtdc3RyaW5ne30sCgl9CgoJZm9yIF8sIG9wdCA6PSByYW5nZSBvcHRzIHsKCQlvcHQoYy5jbGllbnRPcHRzKQoJfQoKCWlmIGMuY2xpZW50T3B0cy5ucyA9PSBuaWwgewoJCWMuY2xpZW50T3B0cy5ucyA9ICZCcm9hZGNhc3ROU3t9Cgl9CgljLm5zID0gYy5jbGllbnRPcHRzLm5zCgoJYy5mdywgZXJyID0gY2hvcmlhLk5ldyhjLmNsaWVudE9wdHMuY2ZnRmlsZSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBpbml0aWFsaXplIENob3JpYTogJXMiLCBlcnIpCgl9CgoJYy5jZmcgPSBjLmZ3LkNvbmZpZ3VyYXRpb24oKQoKCWlmIGMuY2xpZW50T3B0cy5sb2dnZXIgPT0gbmlsIHsKCQljLmNsaWVudE9wdHMubG9nZ2VyID0gYy5mdy5Mb2dnZXIoInt7IC5EREwuTWV0YWRhdGEuTmFtZSB9fSIpCgl9IGVsc2UgewoJCWMuZncuU2V0TG9nZ2VyKGMuY2xpZW50T3B0cy5sb2dnZXIuTG9nZ2VyKQoJfQoKCWRkbGosIGVyciA6PSBiYXNlNjQuU3RkRW5jb2RpbmcuRGVjb2RlU3RyaW5nKHJhd0RETCkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBwYXJzZSBlbWJlZGRlZCBEREw6ICVzIiwgZXJyKQoJfQoKCWVyciA9IGpzb24uVW5tYXJzaGFsKGRkbGosIGMuZGRsKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIG5pbCwgZm10LkVycm9yZigiY291bGQgbm90IHBhcnNlIGVtYmVkZGVkIERETDogJXMiLCBlcnIpCgl9CgoJcmV0dXJuIGMsIG5pbAp9CgovLyBBZ2VudE1ldGFkYXRhIGlzIHRoZSBhZ2VudCBtZXRhZGF0YSB0aGlzIGNsaWVudCBzdXBwb3J0cwpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIEFnZW50TWV0YWRhdGEoKSAqTWV0YWRhdGEgewoJcmV0dXJuICZNZXRhZGF0YXsKCQlMaWNlbnNlOiAgICAgcC5kZGwuTWV0YWRhdGEuTGljZW5zZSwKCQlBdXRob3I6ICAgICAgcC5kZGwuTWV0YWRhdGEuQXV0aG9yLAoJCVRpbWVvdXQ6ICAgICBwLmRkbC5NZXRhZGF0YS5UaW1lb3V0LAoJCU5hbWU6ICAgICAgICBwLmRkbC5NZXRhZGF0YS5OYW1lLAoJCVZlcnNpb246ICAgICBwLmRkbC5NZXRhZGF0YS5WZXJzaW9uLAoJCVVSTDogICAgICAgICBwLmRkbC5NZXRhZGF0YS5VUkwsCgkJRGVzY3JpcHRpb246IHAuZGRsLk1ldGFkYXRhLkRlc2NyaXB0aW9uLAoJfQp9CgovLyBEaXNjb3Zlck5vZGVzIHBlcmZvcm1zIGEgZGlzY292ZXJ5IHVzaW5nIHRoZSBjb25maWd1cmVkIGZpbHRlciBhbmQgbm9kZSBzb3VyY2UKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBEaXNjb3Zlck5vZGVzKGN0eCBjb250ZXh0LkNvbnRleHQpIChub2RlcyBbXXN0cmluZywgZXJyIGVycm9yKSB7CglwLkxvY2soKQoJZGVmZXIgcC5VbmxvY2soKQoKCXJldHVybiBwLm5zLkRpc2NvdmVyKGN0eCwgcC5mdywgcC5maWx0ZXJzKQp9Cgp7eyByYW5nZSAkaSwgJGFjdGlvbiA6PSAuRERMLkFjdGlvbnMgfX0KLy8ge3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19IHBlcmZvcm1zIHRoZSB7eyAkYWN0aW9uLk5hbWUgfCBUb0xvd2VyIH19IGFjdGlvbgovLwovLyBEZXNjcmlwdGlvbjoge3sgJGFjdGlvbi5EZXNjcmlwdGlvbiB9fQp7ey0gaWYgQ2hvcmlhUmVxdWlyZWRJbnB1dHMgJGFjdGlvbiB9fQovLwovLyBSZXF1aXJlZCBJbnB1dHM6Cnt7LSByYW5nZSAkbmFtZSwgJGlucHV0IDo9IENob3JpYVJlcXVpcmVkSW5wdXRzICRhY3Rpb24gfX0KLy8gICAgLSB7eyAkbmFtZSB9fSAoe3sgJGlucHV0LlR5cGUgfCBDaG9yaWFUeXBlVG9Hb1R5cGUgfX0pIC0ge3sgJGlucHV0LkRlc2NyaXB0aW9uIH19Cnt7LSBlbmQgfX0Ke3stIGVuZCB9fQp7ey0gaWYgQ2hvcmlhT3B0aW9uYWxJbnB1dHMgJGFjdGlvbiB9fQovLwovLyBPcHRpb25hbCBJbnB1dHM6Cnt7LSByYW5nZSAkbmFtZSwgJGlucHV0IDo9IENob3JpYU9wdGlvbmFsSW5wdXRzICRhY3Rpb24gfX0KLy8gICAgLSB7eyAkbmFtZSB9fSAoe3sgJGlucHV0LlR5cGUgfCBDaG9yaWFUeXBlVG9Hb1R5cGUgfX0pIC0ge3sgJGlucHV0LkRlc2NyaXB0aW9uIH19Cnt7LSBlbmQgfX0Ke3stIGVuZCB9fQpmdW5jIChwICp7eyAkLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSB7eyAkYWN0aW9uLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0oe3sgJGFjdGlvbiB8IENob3JpYVJlcXVpcmVkSW5wdXRzVG9GdW5jQXJncyB9fSkgKnt7ICRhY3Rpb24uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RlciB7CglkIDo9ICZ7eyAkYWN0aW9uLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXF1ZXN0ZXJ7CgkJb3V0YzogbmlsLAoJCXI6ICZyZXF1ZXN0ZXJ7CgkJCWFyZ3M6ICAgbWFwW3N0cmluZ11pbnRlcmZhY2V7fXsKe3stIHJhbmdlICRuYW1lLCAkaW5wdXQgOj0gQ2hvcmlhUmVxdWlyZWRJbnB1dHMgJGFjdGlvbiB9fQoJCQkJInt7ICRuYW1lIH19Ijoge3sgJG5hbWUgfX1JLAp7ey0gZW5kIH19CgkJCX0sCgkJCWFjdGlvbjogInt7ICRhY3Rpb24uTmFtZSB8IFRvTG93ZXIgfX0iLAoJCQljbGllbnQ6IHAsCgkJfSwKCX0KCglyZXR1cm4gZAp9Cnt7IGVuZCB9fQo=",
	"ddl":           "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19Cgp2YXIgcmF3RERMID0gInt7IC5SYXdEREwgfCBCYXNlNjRFbmNvZGUgfX0i",
	"discover":      "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImNvbnRleHQiCgkic3luYyIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvY2xpZW50L2Rpc2NvdmVyeS9icm9hZGNhc3QiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3RvY29sIgopCgovLyBCcm9hZGNhc3ROUyBpcyBhIE5vZGVTb3VyY2UgdGhhdCB1c2VzIHRoZSBDaG9yaWEgbmV0d29yayBicm9hZGNhc3QgbWV0aG9kIHRvIGRpc2NvdmVyIG5vZGVzCnR5cGUgQnJvYWRjYXN0TlMgc3RydWN0IHsKCW5vZGVDYWNoZSBbXXN0cmluZwoJZiAgICAgICAgICpwcm90b2NvbC5GaWx0ZXIKCglzeW5jLk11dGV4Cn0KCi8vIFJlc2V0IHJlc2V0cyB0aGUgaW50ZXJuYWwgbm9kZSBjYWNoZQpmdW5jIChiICpCcm9hZGNhc3ROUykgUmVzZXQoKSB7CgliLkxvY2soKQoJZGVmZXIgYi5VbmxvY2soKQoKCWIubm9kZUNhY2hlID0gW11zdHJpbmd7fQp9CgovLyBEaXNjb3ZlciBwZXJmb3JtcyB0aGUgZGlzY292ZXJ5IG9mIG5vZGVzIGFnYWluc3QgdGhlIENob3JpYSBOZXR3b3JrCmZ1bmMgKGIgKkJyb2FkY2FzdE5TKSBEaXNjb3ZlcihjdHggY29udGV4dC5Db250ZXh0LCBmdyBDaG9yaWFGcmFtZXdvcmssIGZpbHRlcnMgW11GaWx0ZXJGdW5jKSAoW11zdHJpbmcsIGVycm9yKSB7CgliLkxvY2soKQoJZGVmZXIgYi5VbmxvY2soKQoKCWNvcGllciA6PSBmdW5jKCkgW11zdHJpbmcgewoJCW91dCA6PSBtYWtlKFtdc3RyaW5nLCBsZW4oYi5ub2RlQ2FjaGUpKQoJCWZvciBpLCBuIDo9IHJhbmdlIGIubm9kZUNhY2hlIHsKCQkJb3V0W2ldID0gbgoJCX0KCgkJcmV0dXJuIG91dAoJfQoKCWlmICEoYi5ub2RlQ2FjaGUgPT0gbmlsIHx8IGxlbihiLm5vZGVDYWNoZSkgPT0gMCkgewoJCXJldHVybiBjb3BpZXIoKSwgbmlsCgl9CgoJZXJyIDo9IGIucGFyc2VGaWx0ZXJzKGZpbHRlcnMpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBlcnIKCX0KCglpZiBiLm5vZGVDYWNoZSA9PSBuaWwgewoJCWIubm9kZUNhY2hlID0gW11zdHJpbmd7fQoJfQoKCWNmZyA6PSBmdy5Db25maWd1cmF0aW9uKCkKCW5vZGVzLCBlcnIgOj0gYnJvYWRjYXN0Lk5ldyhmdykuRGlzY292ZXIoY3R4LCBicm9hZGNhc3QuRmlsdGVyKGIuZiksIGJyb2FkY2FzdC5UaW1lb3V0KHRpbWUuU2Vjb25kKnRpbWUuRHVyYXRpb24oY2ZnLkRpc2NvdmVyeVRpbWVvdXQpKSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBbXXN0cmluZ3t9LCBlcnIKCX0KCgliLm5vZGVDYWNoZSA9IG5vZGVzCgoJcmV0dXJuIGNvcGllcigpLCBuaWwKfQoKZnVuYyAoYiAqQnJvYWRjYXN0TlMpIHBhcnNlRmlsdGVycyhmcyBbXUZpbHRlckZ1bmMpIGVycm9yIHsKCWIuZiA9IHByb3RvY29sLk5ld0ZpbHRlcigpCgoJZm9yIF8sIGYgOj0gcmFuZ2UgZnMgewoJCWVyciA6PSBmKGIuZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJcmV0dXJuIGVycgoJCX0KCX0KCglyZXR1cm4gbmlsCn0K",
	"doc":           "e3sgR2VuZXJhdGVkV2FybmluZyB9fQoKLy8gUGFja2FnZSB7eyAuUGFja2FnZSB9fSBpcyBhbiBBUEkgY2xpZW50IHRvIHRoZSBDaG9yaWEge3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgQ2FwaXRhbGl6ZSB9fSBhZ2VudCBWZXJzaW9uIHt7IC5EREwuTWV0YWRhdGEuVmVyc2lvbiB9fS4KLy8KLy8gQWN0aW9uczoKe3stIHJhbmdlICRpLCAkYWN0aW9uIDo9IC5EREwuQWN0aW9ucyB9fQovLyAgICoge3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19IC0ge3sgJGFjdGlvbi5EZXNjcmlwdGlvbiAtfX0Ke3sgZW5kIH19CnBhY2thZ2Uge3sgLlBhY2thZ2UgfX0KCg==",
	"initoptions":   "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImdpdGh1Yi5jb20vc2lydXBzZW4vbG9ncnVzIgopCgp0eXBlIGluaXRPcHRpb25zIHN0cnVjdCB7CgljZmdGaWxlIHN0cmluZwoJbG9nZ2VyICAqbG9ncnVzLkVudHJ5CglucyBOb2RlU291cmNlCglwcm9ncmVzcyBib29sCn0KCi8vIEluaXRpYWxpemF0aW9uT3B0aW9uIGlzIGFuIG9wdGlvbmFsIHNldHRpbmcgdXNlZCB0byBpbml0aWFsaXplIHRoZSBjbGllbnQKdHlwZSBJbml0aWFsaXphdGlvbk9wdGlvbiBmdW5jKG9wdHMgKmluaXRPcHRpb25zKQoKLy8gQ29uZmlnRmlsZSBzZXRzIHRoZSBjb25maWcgZmlsZSB0byB1c2UsIHdoZW4gbm90IHNldCB3aWxsIHVzZSB0aGUgdXNlciBkZWZhdWx0CmZ1bmMgQ29uZmlnRmlsZShmIHN0cmluZykgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLmNmZ0ZpbGUgPSBmCgl9Cn0KCi8vIExvZ2dlciBzZXRzIHRoZSBsb2dnZXIgdG8gdXNlIGVsc2Ugb25lIGlzIG1hZGUgdmlhIHRoZSBDaG9yaWEgZnJhbWV3b3JrCmZ1bmMgTG9nZ2VyKGwgKmxvZ3J1cy5FbnRyeSkgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLmxvZ2dlciA9IGwKCX0KfQoKLy8gRGlzY292ZXJ5IHNldHMgdGhlIE5vZGVTb3VyY2UgdG8gdXNlIHdoZW4gZmluZGluZyBub2RlcyB0byBtYW5hZ2UKZnVuYyBEaXNjb3ZlcnkobnMgTm9kZVNvdXJjZSkgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLm5zID0gbnMKCX0KfQoKLy8gUHJvZ3Jlc3MgZW5hYmxlcyBkaXNwbGF5aW5nIGEgcHJvZ3Jlc3MgYmFyCmZ1bmMgUHJvZ3Jlc3MoKSBJbml0aWFsaXphdGlvbk9wdGlvbiB7CglyZXR1cm4gZnVuYyhvICppbml0T3B0aW9ucykgewoJCW8ucHJvZ3Jlc3MgPSB0cnVlCgl9Cn0K",
	"logging":       "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgpmdW5jIChjICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIGluZm9mKG1zZyBzdHJpbmcsIGEgLi4uaW50ZXJmYWNle30pIHsKCWMuY2xpZW50T3B0cy5sb2dnZXIuSW5mb2YobXNnLCBhLi4uKQp9CgpmdW5jIChjICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIGVycm9yZihtc2cgc3RyaW5nLCBhIC4uLmludGVyZmFjZXt9KSB7CgljLmNsaWVudE9wdHMubG9nZ2VyLkVycm9yZihtc2csIGEuLi4pCn0K",
	"requester":     "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImNvbnRleHQiCgkiZm10IgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvcHJvdG9jb2wiCglycGNjbGllbnQgImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjL2NsaWVudCIKCSJnaXRodWIuY29tL2dvc3VyaS91aXByb2dyZXNzIgoJImdpdGh1Yi5jb20vZmF0aWgvY29sb3IiCikKCi8vIHJlcXVlc3RlciBpcyBhIGdlbmVyaWMgcmVxdWVzdCBoYW5kbGVyCnR5cGUgcmVxdWVzdGVyIHN0cnVjdCB7CgljbGllbnQgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudAoJYWN0aW9uIHN0cmluZwoJYXJncyAgIG1hcFtzdHJpbmddaW50ZXJmYWNle30KCXByb2dyZXNzICp1aXByb2dyZXNzLkJhcgp9CgovLyBkbyBwZXJmb3JtcyB0aGUgcmVxdWVzdApmdW5jIChyICpyZXF1ZXN0ZXIpIGRvKGN0eCBjb250ZXh0LkNvbnRleHQsIGhhbmRsZXIgZnVuYyhwciBwcm90b2NvbC5SZXBseSwgciAqcnBjY2xpZW50LlJQQ1JlcGx5KSkgKCpycGNjbGllbnQuU3RhdHMsIGVycm9yKSB7Cgl0YXJnZXRzIDo9IG1ha2UoW11zdHJpbmcsIGxlbihyLmNsaWVudC50YXJnZXRzKSkKCXZhciBlcnIgZXJyb3IKCglyLmNsaWVudC5Mb2NrKCkKCWNvcHkodGFyZ2V0cywgci5jbGllbnQudGFyZ2V0cykKCWRpc2NvdmVyZXIgOj0gci5jbGllbnQubnMKCWZpbHRlcnMgOj0gci5jbGllbnQuZmlsdGVycwoJZncgOj0gci5jbGllbnQuZncKCXByb2dyZXNzIDo9IHIuY2xpZW50LmNsaWVudE9wdHMucHJvZ3Jlc3MKCglpZiBsZW4odGFyZ2V0cykgPT0gMCB7CgkJaWYgcHJvZ3Jlc3MgewoJCQlmbXQuUHJpbnQoIkRpc2NvdmVyaW5nIG5vZGVzIC4uLi4gIikKCQl9IGVsc2UgewoJCQlyLmNsaWVudC5pbmZvZigiU3RhcnRpbmcgZGlzY292ZXJ5IikKCQl9CgoJCXRhcmdldHMsIGVyciA9IGRpc2NvdmVyZXIuRGlzY292ZXIoY3R4LCBmdywgZmlsdGVycykKCQlpZiBlcnIgIT0gbmlsIHsKCQkJcmV0dXJuIG5pbCwgZXJyCgkJfQoKCQlpZiBsZW4odGFyZ2V0cykgPT0gMCB7CgkJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoIm5vIG5vZGVzIHdlcmUgZGlzY292ZXJlZCIpCgkJfQoKCQlpZiBwcm9ncmVzcyB7CgkJCWZtdC5QcmludGYoIiVkXG5cbiIsIGxlbih0YXJnZXRzKSkKCQl9IGVsc2UgewoJCQlyLmNsaWVudC5pbmZvZigiRGlzY292ZXJlZCAlZCBub2RlcyIsIGxlbih0YXJnZXRzKSkKCQl9Cgl9CgoJb3B0cyA6PSBbXXJwY2NsaWVudC5SZXF1ZXN0T3B0aW9ue3JwY2NsaWVudC5UYXJnZXRzKHRhcmdldHMpfQoJb3B0cyA9IGFwcGVuZChvcHRzLCByLmNsaWVudC5jbGllbnRSUENPcHRzLi4uKQoJci5jbGllbnQuVW5sb2NrKCkKCiAgICAgICAgaWYgcHJvZ3Jlc3MgewogICAgICAgICAgICAgICAgci5jb25maWd1cmVQcm9ncmVzcyhsZW4odGFyZ2V0cykpCiAgICAgICAgfQoKCWFnZW50LCBlcnIgOj0gcnBjY2xpZW50Lk5ldyhyLmNsaWVudC5mdywgci5jbGllbnQuZGRsLk1ldGFkYXRhLk5hbWUsIHJwY2NsaWVudC5EREwoci5jbGllbnQuZGRsKSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBjcmVhdGUgY2xpZW50OiAlcyIsIGVycikKCX0KCglpZiBwcm9ncmVzcyB7CgkJb3B0cyA9IGFwcGVuZChvcHRzLCBycGNjbGllbnQuUmVwbHlIYW5kbGVyKGZ1bmMocHIgcHJvdG9jb2wuUmVwbHksIHJwY3IgKnJwY2NsaWVudC5SUENSZXBseSkgewoJCQlyLnByb2dyZXNzLkluY3IoKQoJCQloYW5kbGVyKHByLCBycGNyKQoJCX0pKQoJfSBlbHNlIHsKCQlvcHRzID0gYXBwZW5kKG9wdHMsIHJwY2NsaWVudC5SZXBseUhhbmRsZXIoaGFuZGxlcikpCgl9CgoJaWYgIXByb2dyZXNzIHsKCQlyLmNsaWVudC5pbmZvZigiSW52b2tpbmcgJXMjJXMgYWN0aW9uIHdpdGggJSN2Iiwgci5jbGllbnQuZGRsLk1ldGFkYXRhLk5hbWUsIHIuYWN0aW9uLCByLmFyZ3MpCgl9CgoJci5jbGllbnQuaW5mb2YoIkludm9raW5nICVzIyVzIGFjdGlvbiB3aXRoICUjdiIsIHIuY2xpZW50LmRkbC5NZXRhZGF0YS5OYW1lLCByLmFjdGlvbiwgci5hcmdzKQoKCXJlcywgZXJyIDo9IGFnZW50LkRvKGN0eCwgci5hY3Rpb24sIHIuYXJncywgb3B0cy4uLikKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBwZXJmb3JtIGRpc2FibGUgcmVxdWVzdDogJXMiLCBlcnIpCgl9CgogICAgICAgIGlmIHByb2dyZXNzIHsKICAgICAgICAgICAgICAgIHVpcHJvZ3Jlc3MuU3RvcCgpCiAgICAgICAgICAgICAgICBmbXQuUHJpbnRsbigpCiAgICAgICAgfQoKCXJldHVybiByZXMuU3RhdHMoKSwgbmlsCn0KCmZ1bmMgKHIgKnJlcXVlc3RlcikgY29uZmlndXJlUHJvZ3Jlc3MoY291bnQgaW50KSB7CiAgICAgICAgaWYgIXIuY2xpZW50LmNsaWVudE9wdHMucHJvZ3Jlc3MgewogICAgICAgICAgICAgICAgcmV0dXJuCiAgICAgICAgfQoKICAgICAgICByLnByb2dyZXNzID0gdWlwcm9ncmVzcy5BZGRCYXIoY291bnQpLkFwcGVuZENvbXBsZXRlZCgpLlByZXBlbmRFbGFwc2VkKCkKICAgICAgICByLnByb2dyZXNzLlByZXBlbmRGdW5jKGZ1bmMoYiAqdWlwcm9ncmVzcy5CYXIpIHN0cmluZyB7CiAgICAgICAgICAgICAgICBpZiBiLkN1cnJlbnQoKSA8IGNvdW50IHsKICAgICAgICAgICAgICAgICAgICAgICAgcmV0dXJuIGNvbG9yLlJlZFN0cmluZyhmbXQuU3ByaW50ZigiJWQgLyAlZCIsIGIuQ3VycmVudCgpLCBjb3VudCkpCiAgICAgICAgICAgICAgICB9CgogICAgICAgICAgICAgICAgcmV0dXJuIGNvbG9yLkdyZWVuU3RyaW5nKGZtdC5TcHJpbnRmKCIlZCAvICVkIiwgYi5DdXJyZW50KCksIGNvdW50KSkKICAgICAgICB9KQoKICAgICAgICB1aXByb2dyZXNzLlN0YXJ0KCkKfQo=",
	"resultdetails": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJInRpbWUiCgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjIgopCgovLyBSZXN1bHREZXRhaWxzIGlzIHRoZSBkZXRhaWxzIGFib3V0IGEgcmVzdWx0CnR5cGUgUmVzdWx0RGV0YWlscyBzdHJ1Y3QgewoJc2VuZGVyICBzdHJpbmcKCWNvZGUgICAgaW50CgltZXNzYWdlIHN0cmluZwoJdHMgICAgICB0aW1lLlRpbWUKfQoKLy8gU3RhdHVzQ29kZSBpcyBhIHJlcGx5IHN0YXR1cyBhcyBkZWZpbmVkIGJ5IE1Db2xsZWN0aXZlIFNpbXBsZVJQQyAtIGludGVnZXJzIDAgdG8gNQovLwovLyBTZWUgdGhlIGNvbnN0YW50cyBPSywgUlBDQWJvcnRlZCwgVW5rbm93blJQQ0FjdGlvbiwgTWlzc2luZ1JQQ0RhdGEsIEludmFsaWRSUENEYXRhIGFuZCBVbmtub3duUlBDRXJyb3IKdHlwZSBTdGF0dXNDb2RlIHVpbnQ4Cgpjb25zdCAoCgkvLyBPSyBpcyB0aGUgcmVwbHkgc3RhdHVzIHdoZW4gYWxsIHdvcmtlZAoJT0sgPSBTdGF0dXNDb2RlKGlvdGEpCgoJLy8gQWJvcnRlZCBpcyBzdGF0dXMgZm9yIHdoZW4gdGhlIGFjdGlvbiBjb3VsZCBub3QgcnVuLCBtb3N0IGZhaWx1cmVzIGluIGFuIGFjdGlvbiBzaG91bGQgc2V0IHRoaXMKCUFib3J0ZWQKCgkvLyBVbmtub3duQWN0aW9uIGlzIHRoZSBzdGF0dXMgZm9yIHVua25vd24gYWN0aW9ucyByZXF1ZXN0ZWQKCVVua25vd25BY3Rpb24KCgkvLyBNaXNzaW5nRGF0YSBpcyB0aGUgc3RhdHVzIGZvciBtaXNzaW5nIGlucHV0IGRhdGEKCU1pc3NpbmdEYXRhCgoJLy8gSW52YWxpZERhdGEgaXMgdGhlIHN0YXR1cyBmb3IgaW52YWxpZCBpbnB1dCBkYXRhCglJbnZhbGlkRGF0YQoKCS8vIFVua25vd25FcnJvciBpcyB0aGUgc3RhdHVzIGdlbmVyYWwgZmFpbHVyZXMgaW4gYWdlbnRzIHNob3VsZCBzZXQgd2hlbiB0aGluZ3MgZ28gYmFkCglVbmtub3duRXJyb3IKKQoKLy8gU2VuZGVyIGlzIHRoZSBpZGVudGl0eSBvZiB0aGUgcmVtb3RlIHRoYXQgcHJvZHVjZWQgdGhlIG1lc3NhZ2UKZnVuYyAoZCAqUmVzdWx0RGV0YWlscykgU2VuZGVyKCkgc3RyaW5nIHsKCXJldHVybiBkLnNlbmRlcgp9CgovLyBPSyBkZXRlcm1pbmVzIGlmIHRoZSByZXF1ZXN0IHdhcyBzdWNjZXNzZnVsCmZ1bmMgKGQgKlJlc3VsdERldGFpbHMpIE9LKCkgYm9vbCB7CglyZXR1cm4gbWNvcnBjLlN0YXR1c0NvZGUoZC5jb2RlKSA9PSBtY29ycGMuT0sKfQoKLy8gU3RhdHVzTWVzc2FnZSBpcyB0aGUgc3RhdHVzIG1lc3NhZ2UgcHJvZHVjZWQgYnkgdGhlIHJlbW90ZQpmdW5jIChkICpSZXN1bHREZXRhaWxzKSBTdGF0dXNNZXNzYWdlKCkgc3RyaW5nIHsKCXJldHVybiBkLm1lc3NhZ2UKfQoKLy8gU3RhdHVzQ29kZSBpcyB0aGUgc3RhdHVzIGNvZGUgcHJvZHVjZWQgYnkgdGhlIHJlbW90ZQpmdW5jIChkICpSZXN1bHREZXRhaWxzKSBTdGF0dXNDb2RlKCkgU3RhdHVzQ29kZSB7CglyZXR1cm4gU3RhdHVzQ29kZShkLmNvZGUpCn0K",
	"rpcoptions":    "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJInRpbWUiCgoJY29yZWNsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL2NsaWVudC9jbGllbnQiCglycGNjbGllbnQgImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjL2NsaWVudCIKKQoKLy8gT3B0aW9uUmVzZXQgcmVzZXRzIHRoZSBjbGllbnQgb3B0aW9ucyB0byB1c2UgYWNyb3NzIHJlcXVlc3RzIHRvIGFuIGVtcHR5IGxpc3QKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25SZXNldCgpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5Mb2NrKCkKCWRlZmVyIHAuVW5sb2NrKCkKCglwLmNsaWVudFJQQ09wdHMgPSBbXXJwY2NsaWVudC5SZXF1ZXN0T3B0aW9ue30KCXAubnMgPSBwLmNsaWVudE9wdHMubnMKCXAudGFyZ2V0cyA9IFtdc3RyaW5ne30KCXAuZmlsdGVycyA9IFtdRmlsdGVyRnVuY3sKCSAgICBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuQWdlbnRGaWx0ZXIoInt7IC5EREwuTWV0YWRhdGEuTmFtZSB9fSIpKSwKCX0KCglyZXR1cm4gcAp9CgovLyBPcHRpb25JZGVudGl0eUZpbHRlciBhZGRzIGFuIGlkZW50aXR5IGZpbHRlcgpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbklkZW50aXR5RmlsdGVyKGYgLi4uc3RyaW5nKSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuTG9jaygpCglkZWZlciBwLlVubG9jaygpCgoJZm9yIF8sIGkgOj0gcmFuZ2UgZiB7CgkJcC5maWx0ZXJzID0gYXBwZW5kKHAuZmlsdGVycywgRmlsdGVyRnVuYyhjb3JlY2xpZW50LklkZW50aXR5RmlsdGVyKGkpKSkKCX0KCglwLm5zLlJlc2V0KCkKCglyZXR1cm4gcAp9CgovLyBPcHRpb25DbGFzc0ZpbHRlciBhZGRzIGEgY2xhc3MgZmlsdGVyCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uQ2xhc3NGaWx0ZXIoZiAuLi5zdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5Mb2NrKCkKCWRlZmVyIHAuVW5sb2NrKCkKCglmb3IgXywgaSA6PSByYW5nZSBmIHsKCQlwLmZpbHRlcnMgPSBhcHBlbmQocC5maWx0ZXJzLCBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuQ2xhc3NGaWx0ZXIoaSkpKQoJfQoKCXAubnMuUmVzZXQoKQoKCXJldHVybiBwCn0KCi8vIE9wdGlvbkZhY3RGaWx0ZXIgYWRkcyBhIGZhY3QgZmlsdGVyCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uRmFjdEZpbHRlcihmIC4uLnN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLkxvY2soKQoJZGVmZXIgcC5VbmxvY2soKQoKCWZvciBfLCBpIDo9IHJhbmdlIGYgewoJCXAuZmlsdGVycyA9IGFwcGVuZChwLmZpbHRlcnMsIEZpbHRlckZ1bmMoY29yZWNsaWVudC5GYWN0RmlsdGVyKGkpKSkKCX0KCglwLm5zLlJlc2V0KCkKCglyZXR1cm4gcAp9CgovLyBPcHRpb25Db21iaW5lZEZpbHRlciBhZGRzIGEgY29tYmluZWQgZmlsdGVyCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uQ29tYmluZWRGaWx0ZXIoZiAuLi5zdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5Mb2NrKCkKCWRlZmVyIHAuVW5sb2NrKCkKCglmb3IgXywgaSA6PSByYW5nZSBmIHsKCQlwLmZpbHRlcnMgPSBhcHBlbmQocC5maWx0ZXJzLCBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuQ29tYmluZWRGaWx0ZXIoaSkpKQoJfQoKCXAubnMuUmVzZXQoKQoKCXJldHVybiBwCn0KCi8vIE9wdGlvbkNvbGxlY3RpdmUgc2V0cyB0aGUgY29sbGVjdGl2ZSB0byB0YXJnZXQKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25Db2xsZWN0aXZlKGMgc3RyaW5nKSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuTG9jaygpCglkZWZlciBwLlVubG9jaygpCgoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkNvbGxlY3RpdmUoYykpCglyZXR1cm4gcAp9CgovLyBPcHRpb25JbkJhdGNoZXMgcGVyZm9ybXMgcmVxdWVzdHMgaW4gYmF0Y2hlcwpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkluQmF0Y2hlcyhzaXplIGludCwgc2xlZXAgaW50KSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuTG9jaygpCglkZWZlciBwLlVubG9jaygpCgoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkluQmF0Y2hlcyhzaXplLCBzbGVlcCkpCglyZXR1cm4gcAp9CgovLyBPcHRpb25EaXNjb3ZlcnlUaW1lb3V0IGNvbmZpZ3VyZXMgdGhlIHJlcXVlc3QgZGlzY292ZXJ5IHRpbWVvdXQsIGRlZmF1bHRzIHRvIGNvbmZpZ3VyZWQgZGlzY292ZXJ5IHRpbWVvdXQKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25EaXNjb3ZlcnlUaW1lb3V0KHQgdGltZS5EdXJhdGlvbikgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLkxvY2soKQoJZGVmZXIgcC5VbmxvY2soKQoKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5EaXNjb3ZlcnlUaW1lb3V0KHQpKQoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uTGltaXRNZXRob2QgY29uZmlndXJlcyB0aGUgbWV0aG9kIHRvIHVzZSB3aGVuIGxpbWl0aW5nIHRhcmdldHMgLSAicmFuZG9tIiBvciAiZmlyc3QiCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uTGltaXRNZXRob2QobSBzdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5Mb2NrKCkKCWRlZmVyIHAuVW5sb2NrKCkKCglwLmNsaWVudFJQQ09wdHMgPSBhcHBlbmQocC5jbGllbnRSUENPcHRzLCBycGNjbGllbnQuTGltaXRNZXRob2QobSkpCglyZXR1cm4gcAp9CgovLyBPcHRpb25MaW1pdFNpemUgc2V0cyBsaW1pdHMgb24gdGhlIHRhcmdldHMsIGVpdGhlciBhIG51bWJlciBvZiBhIHBlcmNlbnRhZ2UgbGlrZSAiMTAlIgpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkxpbWl0U2l6ZShzIHN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLkxvY2soKQoJZGVmZXIgcC5VbmxvY2soKQoKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5MaW1pdFNpemUocykpCglyZXR1cm4gcAp9CgovLyBPcHRpb25MaW1pdFNlZWQgc2V0cyB0aGUgcmFuZG9tIHNlZWQgdXNlZCB0byBzZWxlY3QgdGFyZ2V0cyB3aGVuIGxpbWl0aW5nIGFuZCBsaW1pdCBtZXRob2QgaXMgInJhbmRvbSIKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25MaW1pdFNlZWQocyBpbnQ2NCkgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLkxvY2soKQoJZGVmZXIgcC5VbmxvY2soKQoKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5MaW1pdFNlZWQocykpCglyZXR1cm4gcAp9CgovLyBPcHRpb25UYXJnZXRzIHNldHMgc3BlY2lmaWMgbm9kZSB0YXJnZXRzIHdoaWNoIHdvdWxkIGF2b2lkIGRpc2NvdmVyeSBmb3IgYWxsIGFjdGlvbiBjYWxscyB1bnRpbCByZXNldApmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvblRhcmdldHModCBbXXN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLkxvY2soKQoJZGVmZXIgcC5VbmxvY2soKQoKCXAudGFyZ2V0cyA9IHQKCXJldHVybiBwCn0K",
}
