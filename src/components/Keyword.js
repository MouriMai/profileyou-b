import React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Keyword = () => {
    const [keyword, setKeyword] = useState({});
    let { id } = useParams();

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`/keyword/${id}`, requestOptions)
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
            <h2>Keyword: {keyword.word}</h2>
            <small><em>{keyword.description}</em></small><br />
            <hr />

            {keyword.image_url !== "" &&
                <div className="mb-3">
                    <img src={`${keyword.image_url}`} alt="picture" />
                </div>
            }

            {/* <p>{movie.description}</p> */}
        </div>
    )
}

export default Keyword;