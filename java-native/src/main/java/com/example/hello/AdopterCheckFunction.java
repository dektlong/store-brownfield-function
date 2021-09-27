package com.example.hello;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.client.RestTemplate;
import java.util.concurrent.TimeUnit;

import java.util.function.Function;

@SpringBootApplication
public class AdopterCheckFunction {

    @Value("${TARGET:from-function}")
    String target;

    public static void main(String[] args) {
        SpringApplication.run(AdopterCheckFunction.class, args);
    }

    @Bean
    public Function<String, String> hello() {
        return (in) -> {
		
            String API_CALL = "sample-api";
            
            RestTemplate restTemplate = new RestTemplate();

            try
	    {
   		String apiResults = restTemplate.getForObject(API_CALL, String.class);
	    }

	    catch (Exception e) {
		    /*check failure*/
	    }

            
            return apiResults;
        };
    }
}
