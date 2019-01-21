package main

import "github.com/EngoEngine/glm"

type snake struct {
	bodyparts []glm.Vec2
	direction glm.Vec2
	speed float32
	moveTimer float32
}

func (this *snake) MoveUp() {
	this.direction = glm.Vec2{0.0, -1.0}
}

func (this *snake) MoveDown() {
	this.direction = glm.Vec2{0.0, 1.0}
}

func (this *snake) MoveLeft() {
	this.direction = glm.Vec2{-1.0, 0.0}
}

func (this *snake) MoveRight() {
	this.direction = glm.Vec2{1.0, 0.0}
}

func (this *snake) AddBodypart() {
	var lastBodypart = this.bodyparts[0]
	var newBodypart glm.Vec2 = lastBodypart
	this.bodyparts = append([]glm.Vec2{newBodypart}, this.bodyparts ...)
}

func (this *snake) checkSelfCollision() bool {
	var totalBodyParts int = len(this.bodyparts)
	var head = this.bodyparts[totalBodyParts-1]
	for i := 0; i < totalBodyParts-1; i++{
		var body glm.Vec2 = this.bodyparts[i]
		if head.Equal(&body) {
			return true
		}
	}
	return false
}

func (this *snake) checkPointCollision(point glm.Vec2) bool {
	var totalBodyParts int = len(this.bodyparts)
	var head = this.bodyparts[totalBodyParts-1]
	if head.Equal(&point) {
		return true
	}
	return false
}

func (this *snake) Update() {
	if this.moveTimer >= 1 {
		this.moveTimer = 0.0
		var totalBodyParts int = len(this.bodyparts)
		for i := 0; i < totalBodyParts-1; i++ {
			this.bodyparts[i] = this.bodyparts[i+1]
		}
		this.bodyparts[totalBodyParts-1].AddWith(&this.direction)
	}
	if this.direction.Equal(&glm.Vec2{0.0, 0.0}){
		return
	}
	this.moveTimer += this.speed
}

func (this *snake) SetHeadPosition(position glm.Vec2) {
	var totalBodyParts int = len(this.bodyparts)
	this.bodyparts[totalBodyParts-1] = position
}

func (this *snake) GetHead() glm.Vec2 {
	var totalBodyParts int = len(this.bodyparts)
	return this.bodyparts[totalBodyParts-1]
}

func NewSnake(position glm.Vec2) *snake {
	var bodyparts = make([]glm.Vec2, 0)
	var direction = glm.Vec2{0.0, 0.0}
	var head = position
	bodyparts = append(bodyparts, head)
	var body = head
	body.AddWith(&glm.Vec2{0.0, -1.0})
	bodyparts = append(bodyparts, body)

	return &snake{bodyparts, direction, 0.001, 0.0}
}
