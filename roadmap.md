# Upcoming Releases

__Features & Fixes__

## 0.10.8-alpha (probably)

This version relaxes some validation rules for Event Markdown which imposed unnecessary limits on what kind of event stormings can be used as input for code generation:

* In a real life event storming, an event can occur more than once in a workflow. This is now valid EMD.
* In real life, a command can result in more than one event. And so it is now in EMD.
* Read models can have derived or renamed properties which do not correspond to properties of events.
