embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"4X4\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "4X4_board"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"dropzone\"\n"
  "mask: \"cursor\"\n"
  "mask: \"cards\"\n"
  ""
}
