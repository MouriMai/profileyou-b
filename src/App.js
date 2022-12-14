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
import React from "react";
import { Routes, Route } from "react-router-dom";
import Keyword from "./components/Keyword";
import { Outlet } from "react-router-dom";


const App = () => {
  const [image, updateImage] = useState();
  const [word, updateWord] = useState();
  const [loading, updateLoading] = useState();
  const [keywords, setKeywords] = useState([]);
  const [keyword, setKeyword] = useState({
    id: 0,
    word: "",
    description: "",
    imageUrl: "",
    keywordId: ""
  });


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
        
        fetch(url, requestOptions)
          .then((response) => response.json())
          .then((data) => {
            if (data.error) {
              console.log(data.error);
            } else {
              setKeywords((keywords) => keywords.filter((keyword) => keyword.ID !== id));
              console.log(`${data.Word} Deleted`);
              // navigate("/keywords");
            }
          })
          .catch(err => {
            console.log(err);
          });
      }
    });
  }
    // FINISH DELETE

    const handleSubmit = (e) => {
      e.preventDefault();
      console.log({ word });
      generate(word)
    };

  const generate = async (word) => {
    updateLoading(true);
    // const request = await axios.post(`http://localhost:8080/keyword/create/${word}`);
    // const result = await axios.get(`http://localhost:8080/keyword/create/${word}`);
    // updateImage(result.data);
    const requestBody = word;
    
    let headers = new Headers();
    // headers.append("Content-Type", "application/json");
        // headers.append("Authorization", "Bearer " + jwtToken)
        const requestOptions = {
          method: "POST",
          headers: headers,
          body: JSON.stringify(requestBody)
        };

        const url = `http://localhost:8080/keyword/create/${word}`;
        fetch(url, requestOptions)
          .then((response) => response.json())
          .then((data) => {
            console.log(data);
            if (data.error) {
              console.log(data.error);
            } else {
              setKeywords(data);
            }
          })
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
          <form method="post" onSubmit={handleSubmit}>
            <Input
              id="word"
              value={word}
              name="word"
              onChange={(e) => updateWord(e.target.value)}
              width={"350px"}
            ></Input>
            <button type="submit">Send</button>
            <Button colorScheme={"yellow"}>
              Generate
            </Button>
          </form>
        </Wrap>

        {loading ? (
          <Stack>
            <SkeletonCircle />
            <SkeletonText />
          </Stack>
        ) : image ? (
          <Image src={`data:image/png;base64,${image}`} boxShadow="lg" />
        ) : null}

        <pre>{JSON.stringify(word)}</pre>
        {/* Outletは共通NavBarとかを望むとき */}
        {/* <Outlet context={{keywords, confirmDelete}}/> */}
        {/* Routeの一部にしない下記の記述は居座るから🆖 */}
        {/* <Keywords keywords={keywords} confirmDelete={confirmDelete}/> */}
        {/* このルーティングめちゃくちゃ苦労した　何だこれ */}
        <Routes>
          <Route path={`/keywords`} element={<Keywords keywords={keywords} confirmDelete={confirmDelete}/>} />
          <Route path={`/keywords/:id`} element={<Keyword />} />
        </Routes>

      </Container>
    </ChakraProvider>
  );
}

export default App;

