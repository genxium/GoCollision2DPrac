package main

import (
	"fmt"
  "github.com/ByteArena/box2d"
)

func observeContact(world *box2d.B2World, itContacts box2d.B2ContactInterface) {
  fmt.Printf("Got an AABB contact!\n")
  fixtureA := itContacts.GetFixtureA()
  shapeA := fixtureA.GetShape()
  fixtureB := itContacts.GetFixtureB()
  shapeB := fixtureB.GetShape()
  if itContacts.IsTouching() {
    fmt.Printf("Got a touching contact with shapeA %v and shapeB %v!\n", shapeA, shapeB)
  }
}

func moveDynamicBody(body *box2d.B2Body, toTargetPosition box2d.B2Vec2, inSeconds float64) {
  if body.GetType() != box2d.B2BodyType.B2_dynamicBody {
    fmt.Printf("This is NOT a dynamic body!\n")
    return
  }
  /*
  currentPosition := body.GetPosition()
  positionDiff := box2d.B2Vec2Sub(toTargetPosition, currentPosition)
  */
  body.SetTransform(toTargetPosition, 0.0)
  body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
  body.SetAngularVelocity(0.0)
}

func main() {
	gravity := box2d.MakeB2Vec2(0.0, 0.0)
	world := box2d.MakeB2World(gravity)

  pos1 := box2d.MakeB2Vec2(100.0, -100.0)
  pos2 := box2d.MakeB2Vec2(100.0, 0.0)
  pos3 := box2d.MakeB2Vec2(100.0, 300.0)

  var polygon *box2d.B2Body
  {
    bodyDef := box2d.MakeB2BodyDef()
    bodyDef.Position.Set(pos1.X, pos1.Y)
    bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
		polygon = world.CreateBody(&bodyDef)

    pointsCount := 5
    polygonCorners := [...]float64{
      0, 100,
      -100, 0,
      -100, -100,
      100, -100,
      100, 0,
    }
		vertices := make([]box2d.B2Vec2, pointsCount)
		for i := 0; i < pointsCount; i++ {
			vertices[i].Set(polygonCorners[2*i], polygonCorners[2*i + 1])
		}

		shape := box2d.MakeB2PolygonShape()
		shape.Set(vertices, pointsCount)

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 0.0
		polygon.CreateFixtureFromDef(&fd)

    fmt.Printf("\n#################\nCreated polygon at %v\n", polygon.GetPosition())
  }

  var circle *box2d.B2Body
  {
    bodyDef := box2d.MakeB2BodyDef()
    bodyDef.Position.Set(100.0, 100.0)
    bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
		circle = world.CreateBody(&bodyDef)

    shape := box2d.MakeB2CircleShape()
    shape.M_radius = 32.0

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 0.0
		circle.CreateFixtureFromDef(&fd)
    fmt.Printf("\n#################\nCreated circle at %v\n", circle.GetPosition())
  }

  uniformTimeStepSeconds := 1.0/60.0
  uniformVelocityIterations := 0
  uniformPositionIterations := 0
  {
    targetPos := pos2
    fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
    moveDynamicBody(polygon, targetPos, uniformTimeStepSeconds)
    world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
    itContacts := world.GetContactList()
    for itContacts != nil {
      observeContact(&world, itContacts)
      itContacts = itContacts.GetNext()
    }
  }
  {
    targetPos := pos3
    fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
    moveDynamicBody(polygon, targetPos, uniformTimeStepSeconds)
    world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
    itContacts := world.GetContactList()
    for itContacts != nil {
      observeContact(&world, itContacts)
      itContacts = itContacts.GetNext()
    }
  }
  {
    targetPos := pos2
    fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
    moveDynamicBody(polygon, targetPos, uniformTimeStepSeconds)
    world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
    itContacts := world.GetContactList()
    for itContacts != nil {
      observeContact(&world, itContacts)
      itContacts = itContacts.GetNext()
    }
  }
}
