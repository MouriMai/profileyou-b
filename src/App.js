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
import Swal from "sweetalert2";
import { useEffect, useState } from "react";
import Keywords from "./components/Keywords";


const App = () => {
  const [image, updateImage] = useState();
  const [prompt, updatePrompt] = useState();
  const [loading, updateLoading] = useState();
  const [keywords, setKeywords] = useState([]);


  // START FETCHING
  useEffect(() => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    }

    fetch(`http://localhost:8080/keywords`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setKeywords(data);
      })
      .catch(err => {
        console.log(err);
      })
  }, []);
  // FINISH FETCHING

  // DELETE
  function confirmDelete(id) {
    const url = `http://localhost:8080/keyword/delete/${id}`;

    Swal.fire({
      title: 'Delete keyword?',
      text: "You cannot undo this action!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
      if (result.isConfirmed) {
        // axios.delete(url)
        // .then(res => {
        //   const keywords = this.state.keywords.filter(keyword => keyword.id !== id);
        //   console.log("Delete from react:");
        //   setKeywords({keywords})
        //   console.log(res.data);
        // })
        let headers = new Headers();
        // headers.append("Authorization", "Bearer " + jwtToken)
        const requestOptions = {
          method: "DELETE",
          headers: headers,
        };

        const url = `http://localhost:8080/keyword/delete/${id}`;
        console.log(url);
        fetch(`/keyword/delete/${id}`, requestOptions)
          .then((response) => response.json())
          .then((data) => {
            if (data.error) {
              console.log("Something's Error");
              console.log(data.error);
            } else {
              console.log("Finish");
              setKeywords(data);
              // navigate("/keywords");
            }
          })
          .catch(err => {
            console.log("Error was caught:");
            console.log(err);
          });
      }
    });
  }
    // FINISH DELETE


  const addKeywordHandler = (keyword) => {
    // event.preventDefault()
    console.log(keywords);
    setKeywords([...keywords, keyword]);
    console.log(keyword);
  };

  const generate = async (prompt) => {
    addKeywordHandler(prompt)
    updateLoading(true);
    const request = await axios.post(`http://localhost:8080/keyword/create/${prompt}`);
    // const result = await axios.get(`http://localhost:8080/keyword/create/${prompt}`);
    console.log(request);
    // console.log(result);
    // updateImage(result.data);
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
          <Button onClick={(e) => {
            e.preventDefault()
            generate(prompt)
            }} colorScheme={"yellow"}>
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

        <Keywords keywords={keywords} confirmDelete={confirmDelete}/>

      </Container>
    </ChakraProvider>
  );
}

export default App;
