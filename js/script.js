// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

function calculatePressed() {
  //This function calculates the volume of a Sphere
  //Input through Textfields
  const radius = parseFloat(document.getElementById("radius").value)

  //Process
  const volume = (4 / 3) * Math.PI * radius ** 3

  //Output
  document.getElementById("answer").innerHTML =
    "The Volume of the Sphere is: " + volume.toFixed(2) + "cm³"
}
