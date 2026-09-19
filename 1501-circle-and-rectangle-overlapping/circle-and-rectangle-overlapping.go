func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    // Find the closest x on the rectangle to the circle's center
    closestX := clamp(xCenter, x1, x2)
    // Find the closest y on the rectangle to the circle's center
    closestY := clamp(yCenter, y1, y2)

    // Calculate the distance squared from the circle's center to this closest point
    dx := xCenter - closestX
    dy := yCenter - closestY
    distSquared := (dx * dx) + (dy * dy)

    // Compare squared distance with squared radius
    return distSquared <= (radius * radius)
}

func clamp(val, min, max int) int {
    if val < min {
        return min
    }
    if val > max {
        return max
    }
    return val
}
