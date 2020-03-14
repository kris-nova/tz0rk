# Scenes

These are YAML (of fucking course they are) files that represent a scene in the game. Feel free to PR scenes and they will randomly be drawn at runtime.

# Example scene

```yaml

name: "Example Name"
enable: false # Turn the scene on/off
message: "280 characters or less that will be the content of the tweet" 
danger: 7 # A value 0-9, where 0 is danger free and 9 is the highest danger
options: 
	-	message: "Use the $inventory"
    		inventoryUse: "goose"
       		useMessage: "/goose"
    	-	message: "Pick up the diamonds"
    		inventoryAdd: "diamonds"
    	-	message: "Go deeper into the cave"
```