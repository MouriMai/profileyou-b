import { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";
import { Link, redirect } from "react-router-dom";
import { Button } from "@chakra-ui/react";
import Swal from "sweetalert2";
import axios from 'axios'

const Keywords = (props) => {
  // const navigate = useNavigate();
  const [keywords, setKeywords] = useState([]);

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

const confirmDelete = (id) => {
  const url = `http://localhost:8080/keyword/delete/${id}`

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
            }

            console.log("Fetch starting!");
            const url = `http://localhost:8080/keyword/delete/${id}`
            console.log(url);
          fetch(`keyword/delete/${id}`, requestOptions)
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
              console.log(err)
            });
        }
      })
  }

  return (
    <div>
      <h2>📘  Keywords History</h2>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Keyword</th>
            <th>Description</th>
            <th>......</th>
          </tr>
        </thead>
        <tbody>
          {keywords.map((k, index) => (
            <tr key={index}>
              <td>{k.Word}
                {/* <Link to={`/keywords/${k.ID}`}>
                                    {k.word}
                                </Link> */}
              </td>
              <td>{k.Description}</td>
              <td>
                {k.ImageUrl !== "" &&
                  <div className="mb-3">
                    {/* <img src={`${k.ImageUrl}`} alt="generated-img" /> */}
                    <img src={`https://res.cloudinary.com/dokzsbu2v/image/upload/v1670479204/development/images_xa8j85.png`} alt="generated-img" />
                  </div>
                }

              </td>
              <td>{k.CreatedAt}</td>
              <td><Button onClick={() => confirmDelete(k.ID)} colorScheme={"gray"}  variant={"outline"}>
            🗑️ Delete
          </Button></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default Keywords;
