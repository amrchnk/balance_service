--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3
-- Dumped by pg_dump version 13.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: balance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.balance (
    id integer NOT NULL,
    user_id integer NOT NULL,
    balance numeric(15,4) NOT NULL,
    CONSTRAINT balance_balance_check CHECK ((balance >= (0)::numeric))
);


ALTER TABLE public.balance OWNER TO postgres;

--
-- Name: balance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.balance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.balance_id_seq OWNER TO postgres;

--
-- Name: balance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.balance_id_seq OWNED BY public.balance.id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    user_id integer,
    type_t character varying(50) NOT NULL,
    amount numeric(15,4) NOT NULL,
    description character varying(100),
    created timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: balance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.balance ALTER COLUMN id SET DEFAULT nextval('public.balance_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Data for Name: balance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.balance (id, user_id, balance) FROM stdin;
7	3	1070.0000
10	4	70.0000
1	1	60.4500
13	2	210.0000
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, user_id, type_t, amount, description, created) FROM stdin;
1	1	increase in balance	400.4500		2021-09-06 22:23:19.134096
2	1	increase in balance	130.0000		2021-09-06 22:42:11.858516
3	1	decrease in balance	130.0000		2021-09-06 22:42:25.041675
4	1	decrease in balance	130.0000		2021-09-06 22:42:27.447356
5	3	increase in balance	500.0000		2021-09-06 23:18:06.213599
6	3	increase in balance	500.0000		2021-09-06 23:18:07.145871
7	3	incoming transfer	70.0000	money transfer from user with id=1	2021-09-06 23:27:25.888326
8	1	outgoing transfer	70.0000	money transfer to user with id=3	2021-09-06 23:27:25.889813
9	4	incoming transfer	70.0000	money transfer from user with id=1	2021-09-06 23:27:39.232136
10	1	outgoing transfer	70.0000	money transfer to user with id=4	2021-09-06 23:27:39.233541
11	1	increase in balance	70.0000		2021-09-06 23:27:58.153518
12	1	increase in balance	70.0000		2021-09-06 23:28:37.715262
13	2	incoming transfer	70.0000	money transfer from user with id=1	2021-09-06 23:29:36.976598
14	1	outgoing transfer	70.0000	money transfer to user with id=2	2021-09-06 23:29:36.978176
15	2	incoming transfer	70.0000	money transfer from user with id=1	2021-09-07 19:09:01.593811
16	1	outgoing transfer	70.0000	money transfer to user with id=2	2021-09-07 19:09:01.612961
17	2	incoming transfer	70.0000	money transfer from user with id=1	2021-09-07 19:09:04.376356
18	1	outgoing transfer	70.0000	money transfer to user with id=2	2021-09-07 19:09:04.377873
\.


--
-- Name: balance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.balance_id_seq', 15, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 18, true);


--
-- Name: balance balance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.balance
    ADD CONSTRAINT balance_pkey PRIMARY KEY (id);


--
-- Name: balance balance_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.balance
    ADD CONSTRAINT balance_user_id_key UNIQUE (user_id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.balance(user_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

