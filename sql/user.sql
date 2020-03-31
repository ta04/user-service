--
-- PostgreSQL database dump
--

-- Dumped from database version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: users; Type: DATABASE; Owner: sleepingnext
--

CREATE DATABASE users;

ALTER DATABASE users OWNER TO sleepingnext;


--
-- Name: users; Type: TABLE; Schema: public; Owner: sleepingnext
--

\connect users;
CREATE TABLE public.users (
    id bigint NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    email_address character varying(255) NOT NULL,
    phone_number character varying(255) NOT NULL,
    date_of_birth date NOT NULL,
    address character varying(1000) NOT NULL,
    role character varying(255) NOT NULL,
    credit_card_number character varying(255) NOT NULL,
    credit_card_type character varying(255) NOT NULL,
    credit_card_expired_month character varying(255) NOT NULL,
    credit_card_expired_year character varying(255) NOT NULL,
    credit_card_cvv character varying(255) NOT NULL,
    status character varying(255) NOT NULL
);


ALTER TABLE public.users OWNER TO "sleepingnext";

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: sleepingnext
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO "sleepingnext";

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sleepingnext
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: sleepingnext
--

COPY public.users (id, first_name, last_name, username, password, email_address, phone_number, date_of_birth, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year, credit_card_cvv, status) FROM stdin;
1	admin	admin	admin	admin	admin@skit.com	081271762836	2020-03-12	Admin's Home	admin	empty	empty	empty	empty	empty	active
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sleepingnext
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

