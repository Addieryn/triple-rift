components {
  id: "card"
  component: "/main/card.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"back\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
}
