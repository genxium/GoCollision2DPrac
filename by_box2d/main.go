package main

import (
	"fmt"
	"github.com/ByteArena/box2d"
)

func prettyPrintFixture(fix *box2d.B2Fixture) {
	fmt.Printf("\t\tfriction:\t%v\n", fix.M_friction)
	fmt.Printf("\t\trestitution:\t%v\n", fix.M_restitution)
	fmt.Printf("\t\tdensity:\t%v\n", fix.M_density)
	fmt.Printf("\t\tisSensor:\t%v\n", fix.M_isSensor)
	fmt.Printf("\t\tfilter.categoryBits:\t%d\n", fix.M_filter.CategoryBits)
	fmt.Printf("\t\tfilter.maskBits:\t%d\n", fix.M_filter.MaskBits)
	fmt.Printf("\t\tfilter.groupIndex:\t%d\n", fix.M_filter.GroupIndex)

	switch fix.M_shape.GetType() {
	case box2d.B2Shape_Type.E_circle:
		{
			s := fix.M_shape.(*box2d.B2CircleShape)
			fmt.Printf("\t\tb2CircleShape shape: {\n")
			fmt.Printf("\t\t\tradius:\t%v\n", s.M_radius)
			fmt.Printf("\t\t\toffset:\t%v\n", s.M_p)
			fmt.Printf("\t\t}\n")
		}
		break

	case box2d.B2Shape_Type.E_polygon:
		{
			s := fix.M_shape.(*box2d.B2PolygonShape)
			fmt.Printf("\t\tb2PolygonShape shape: {\n")
			for i := 0; i < s.M_count; i++ {
				fmt.Printf("\t\t\t%v\n", s.M_vertices[i])
			}
			fmt.Printf("\t\t}\n")
		}
		break

	default:
		break
	}
}

func prettyPrintBody(body *box2d.B2Body) {
	bodyIndex := body.M_islandIndex

	fmt.Printf("{\n")
	fmt.Printf("\ttype:\t%d\n", body.M_type)
	fmt.Printf("\tposition:\t%v\n", body.GetPosition())
	fmt.Printf("\tangle:\t%v\n", body.M_sweep.A)
	fmt.Printf("\tlinearVelocity:\t%v\n", body.GetLinearVelocity())
	fmt.Printf("\tangularVelocity:\t%v\n", body.GetAngularVelocity())
	fmt.Printf("\tlinearDamping:\t%v\n", body.M_linearDamping)
	fmt.Printf("\tangularDamping:\t%v\n", body.M_angularDamping)
	fmt.Printf("\tallowSleep:\t%d\n", body.M_flags&box2d.B2Body_Flags.E_autoSleepFlag)
	fmt.Printf("\tawake:\t%d\n", body.M_flags&box2d.B2Body_Flags.E_awakeFlag)
	fmt.Printf("\tfixedRotation:\t%d\n", body.M_flags&box2d.B2Body_Flags.E_fixedRotationFlag)
	fmt.Printf("\tbullet:\t%d\n", body.M_flags&box2d.B2Body_Flags.E_bulletFlag)
	fmt.Printf("\tactive:\t%d\n", body.M_flags&box2d.B2Body_Flags.E_activeFlag)
	fmt.Printf("\tgravityScale:\t%v\n", body.M_gravityScale)
	fmt.Printf("\tislandIndex:\t%v\n", bodyIndex)
	fmt.Printf("\tfixtures: {\n")
	for f := body.M_fixtureList; f != nil; f = f.M_next {
		prettyPrintFixture(f)
	}
	fmt.Printf("\t}\n")
	fmt.Printf("}\n")
}

func observeContact(world *box2d.B2World, itContacts box2d.B2ContactInterface, toDestroyBody *box2d.B2Body, shouldDestroyTheSpecifiedBodyIfTouching bool) {
	fmt.Printf("Got an AABB contact!\n")
	fixtureA := itContacts.GetFixtureA()
	fixtureB := itContacts.GetFixtureB()
	if itContacts.IsTouching() {
		fmt.Printf("Got a touching contact with\nfixtureA\n")
		prettyPrintFixture(fixtureA)
		fmt.Printf("and\nfixtureB\n")
		prettyPrintFixture(fixtureB)

   if shouldDestroyTheSpecifiedBodyIfTouching {
	   fmt.Printf("Destroying the touching target body\n")
     world.DestroyBody(toDestroyBody)
   }
	}
}

func moveDynamicBody(body *box2d.B2Body, pToTargetPos *box2d.B2Vec2, inSeconds float64) {
	if body.GetType() != box2d.B2BodyType.B2_dynamicBody {
		fmt.Printf("This is NOT a dynamic body!\n")
		return
	}
	body.SetTransform(*pToTargetPos, 0.0)
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
			vertices[i].Set(polygonCorners[2*i], polygonCorners[2*i+1])
		}

		shape := box2d.MakeB2PolygonShape()
		shape.Set(vertices, pointsCount)
		shape.SetRadius(0.0)

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 0.0
		polygon.CreateFixtureFromDef(&fd)

		fmt.Printf("\n#################\nCreated a polygon\n")
	  prettyPrintBody(polygon)
	}

	var circle *box2d.B2Body
	{
		bodyDef := box2d.MakeB2BodyDef()
		bodyDef.Position.Set(100.0, 100.0)
		bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
		circle = world.CreateBody(&bodyDef)

		shape := box2d.MakeB2CircleShape()
		shape.SetRadius(32.0)

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 0.0
		circle.CreateFixtureFromDef(&fd)
		fmt.Printf("\n#################\nCreated a circle\n")
	}

	uniformTimeStepSeconds := 1.0 / 60.0
	uniformVelocityIterations := 0
	uniformPositionIterations := 0
	{
		targetPos := pos2
		fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
		moveDynamicBody(polygon, &targetPos, uniformTimeStepSeconds)
    /**
    * During the immediately following statement, effective contact(s) will be discovered by `world.M_contactManager.FindNewContacts()` and marked as `IsTouching` by `world.M_contactManager.Collide()`.
    */
		world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
		itContacts := world.GetContactList()
		for itContacts != nil {
			observeContact(&world, itContacts, circle, false)
			itContacts = itContacts.GetNext()
		}
	}
	{
		targetPos := pos3
		fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
		moveDynamicBody(polygon, &targetPos, uniformTimeStepSeconds)
    /**
    * During the immediately following statement, the previously effective contact(s) will be removed from `world.M_contactManager`.
    */
		world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
		itContacts := world.GetContactList()
		for itContacts != nil {
			observeContact(&world, itContacts, circle, false)
			itContacts = itContacts.GetNext()
		}
	}
	{
		targetPos := pos2
		fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
		moveDynamicBody(polygon, &targetPos, uniformTimeStepSeconds)
    /**
    * During the immediately following statement,
    * - at the beginning no new fixture has been added since creation of the only two bodies, thus `world.M_contactManager.FindNewContacts` won't be called,  
    * - then there's nothing existing in `world.M_contactManager.M_contactList`, thus `world.M_contactManager.Collide()` won't mark anything as `IsTouching`,
    * - at the end new AABB contacts will be discovered by `world.Solve(...)`.
    */
		world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
    /**
    * During the immediately following statement,
    * - at the beginning no new fixture has been added since creation of the only two bodies, thus `world.M_contactManager.FindNewContacts` won't be called,  
    * - then there's existing AABB contacts in `world.M_contactManager.M_contactList`, thus `world.M_contactManager.Collide()` will mark `IsTouching` appropriately.
    *
    * Such specific order of contact(s) management of each `world.Step(...)` is a characteristic of box2d, which is convenient to make use of in HighFPS applications. 
    */
		world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
		itContacts := world.GetContactList()
		for itContacts != nil {
			observeContact(&world, itContacts, circle, true)
			itContacts = itContacts.GetNext()
		}
	}
  {
		targetPos := pos3
		fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
		moveDynamicBody(polygon, &targetPos, uniformTimeStepSeconds)
		world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
		itContacts := world.GetContactList()
		for itContacts != nil {
			observeContact(&world, itContacts, circle, false)
			itContacts = itContacts.GetNext()
		}
  }
  {
		targetPos := pos2
		fmt.Printf("\n#################\nMoving the polygon to %v\n", targetPos)
		moveDynamicBody(polygon, &targetPos, uniformTimeStepSeconds)
		world.Step(uniformTimeStepSeconds, uniformVelocityIterations, uniformPositionIterations)
		itContacts := world.GetContactList()
		for itContacts != nil {
			observeContact(&world, itContacts, circle, false)
			itContacts = itContacts.GetNext()
		}
  }
}
