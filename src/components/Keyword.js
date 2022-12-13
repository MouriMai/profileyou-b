import React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Keyword = (props) => {
    const [keyword, setKeyword] = useState({});
    let { id } = useParams();

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`/keywords/${id}`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setKeyword(data);
            })
            .catch(err => {
                console.log(err);
            })
    }, [id])


    return(
        <div>
            <h2>Keyword: {keyword.Word}</h2>
            <small><em>{keyword.Description}</em></small><br />
            <hr />

            {keyword.imageUrl !== "" &&
                <div className="mb-3">
                    <img src={`${keyword.ImageUrl}`} alt="picture" />
                </div>
            }

            {/* <p>{movie.description}</p> */}
        </div>
    )
}

export default Keyword;