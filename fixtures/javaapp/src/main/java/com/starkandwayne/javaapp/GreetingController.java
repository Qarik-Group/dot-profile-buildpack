package com.starkandwayne.javaapp;

import java.io.File;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class GreetingController {

    @RequestMapping("/")
    public @ResponseBody String greeting() {
        File f = new File("file-created-by-profiled");
        if(f.exists() && !f.isDirectory()) {
            return "Found created file file-created-by-profiled";
        }
        return "Cannot find file file-created-by-profiled";
    }
}
