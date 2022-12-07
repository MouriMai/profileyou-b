--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-2.pgdg110+2)
-- Dumped by pg_dump version 14.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages =


--
-- Name: keywords; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.keywords (
    id integer NOT NULL,
    word character varying(512),
    description text,
    image_url character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);



--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.keywords ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);




--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.keywords (id, word, description, image_url, created_at, updated_at, deleted_at) FROM stdin;
1	Highlander	This is a man who buys a cup of coffee. 	/8Z8dptJEypuLoOQro1WugD855YE.jpg	2022-09-23 00:00:00	2022-09-23 00:00:00
2	Christmas	Children surrounded by a bunch of presents		/ceG9VzoRAVGwivFU403Wc3AHRys.jpg	2022-09-23 00:00:00	2022-09-23 00:00:00
\.




--
-- PostgreSQL database dump complete
--

