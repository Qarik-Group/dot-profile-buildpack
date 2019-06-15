package com.starkandwayne.javaapp;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class GreetingController {

    @RequestMapping("/")
    public @ResponseBody String greeting() {
        return "Cannot find file";
    }
}
