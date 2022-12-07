import {
  ChakraProvider,
  Heading,
  Container,
  Text,
  Link,
  Wrap,
  Input,
  Stack,
  Button,
  Image,
  SkeletonCircle,
  SkeletonText

}
  from "@chakra-ui/react"
import axios from "axios";
import { useState } from "react";
import Keywords from "./components/Keywords";

const App = () => {
  const [image, updateImage] = useState();
  const [prompt, updatePrompt] = useState();
  const [loading, updateLoading] = useState();

  const generate = async (prompt) => {
    updateLoading(true);
    const request = await axios.post(`http://localhost:8080/keyword/create/${prompt}`);
    const result = await axios.get(`http://localhost:8080/keyword/create/${prompt}`);
    console.log(result, request);
    updateImage(result.data);
    updateLoading(false);
  };

  return (
    <ChakraProvider>
      <Container>
        <Heading>Profile You🚀</Heading>
        <Text marginBottom={"10px"}>
          This application examines the trend of the given word in Twitter to generate images
          using the Dall・E API. More information can be found here{" "}
          <Link href={"#"}>
            Web
          </Link>
        </Text>
        <div className="App">
          Profile You!
        </div>
        <Wrap marginBottom={"10px"}>
          <Input
            value={prompt}
            onChange={(e) => updatePrompt(e.target.value)}
            width={"350px"}
          ></Input>
          <Button onClick={(e) => generate(prompt)} colorScheme={"yellow"}>
            Generate
          </Button>
        </Wrap>

        {loading ? (
          <Stack>
            <SkeletonCircle />
            <SkeletonText />
          </Stack>
        ) : image ? (
          <Image src={`data:image/png;base64,${image}`} boxShadow="lg" />
        ) : null}

        <pre>{JSON.stringify(prompt)}</pre>

        <Keywords />

      </Container>
    </ChakraProvider>
  );
}

export default App;
