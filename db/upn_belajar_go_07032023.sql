--
-- PostgreSQL database dump
--

-- Dumped from database version 10.21 (Ubuntu 10.21-1.pgdg22.04+1)
-- Dumped by pg_dump version 10.21 (Ubuntu 10.21-1.pgdg22.04+1)

-- Started on 2023-03-07 11:41:05 WIB

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

DROP DATABASE upn_belajar_go;
--
-- TOC entry 3011 (class 1262 OID 17021)
-- Name: upn_belajar_go; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE upn_belajar_go WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE upn_belajar_go OWNER TO postgres;

\connect upn_belajar_go

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
-- TOC entry 1 (class 3079 OID 13104)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 3014 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 17022)
-- Name: jenis_mitra; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jenis_mitra (
    id character varying(36) NOT NULL,
    nama_jenis_mitra character varying(200),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.jenis_mitra OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 17044)
-- Name: kelas_siswa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kelas_siswa (
    id character varying(36) NOT NULL,
    id_kelas character varying(36),
    tahun_ajaran text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.kelas_siswa OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 17053)
-- Name: kelas_siswa_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kelas_siswa_detail (
    id character varying(36) NOT NULL,
    id_kelas_siswa character varying(36),
    id_siswa character varying(36),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.kelas_siswa_detail OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 17034)
-- Name: m_kelas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_kelas (
    id character varying(36) NOT NULL,
    kode character varying(100),
    nama character varying(100),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.m_kelas OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 17028)
-- Name: m_siswa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_siswa (
    id character varying(36) NOT NULL,
    nama character varying(100),
    kelas character varying(100),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.m_siswa OWNER TO postgres;

--
-- TOC entry 3001 (class 0 OID 17022)
-- Dependencies: 196
-- Data for Name: jenis_mitra; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.jenis_mitra (id, nama_jenis_mitra, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('67153215-b6cd-44c1-a20e-12321f82e572', 'RISET', '2023-03-07 06:56:41.942254', NULL, '667ff26a-d8a0-4ae1-9a5a-277305f404d1', NULL, false);


--
-- TOC entry 3004 (class 0 OID 17044)
-- Dependencies: 199
-- Data for Name: kelas_siswa; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.kelas_siswa (id, id_kelas, tahun_ajaran, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('4c1905fd-5fa7-42ed-97dd-f29c7cee6e44', '54c9cd72-9541-4dc4-8864-16dc8236cfcc', '2023/2024', '2023-03-07 10:44:15.692476', NULL, '', NULL, false);


--
-- TOC entry 3005 (class 0 OID 17053)
-- Dependencies: 200
-- Data for Name: kelas_siswa_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.kelas_siswa_detail (id, id_kelas_siswa, id_siswa, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('c0f74cdf-16c3-4d92-9538-0c00bf1f5b2a', '4c1905fd-5fa7-42ed-97dd-f29c7cee6e44', '5428ec3a-4a8e-444b-aa16-dd8f9d2cb0ee', '2023-03-07 10:44:15.692478', NULL, '', NULL, false);
INSERT INTO public.kelas_siswa_detail (id, id_kelas_siswa, id_siswa, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('3ab935c2-586f-40e8-85a2-e7aa2d433010', '4c1905fd-5fa7-42ed-97dd-f29c7cee6e44', 'eb6a8af6-a218-4f41-a96e-6b982096bc89', '2023-03-07 10:44:15.692481', NULL, '', NULL, false);


--
-- TOC entry 3003 (class 0 OID 17034)
-- Dependencies: 198
-- Data for Name: m_kelas; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.m_kelas (id, kode, nama, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('54c9cd72-9541-4dc4-8864-16dc8236cfcc', 'X-IPA-B', 'X IPA B', '2023-02-28 08:26:53.492433', '2023-02-28 08:26:53.492798', '', NULL, false);
INSERT INTO public.m_kelas (id, kode, nama, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('279db6e8-6f00-4f20-8c22-a5ce914c878b', 'X-IPA-A', 'X IPA A', '2023-02-28 08:25:41.270441', '2023-02-28 08:30:36.543048', '', '', false);


--
-- TOC entry 3002 (class 0 OID 17028)
-- Dependencies: 197
-- Data for Name: m_siswa; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('5428ec3a-4a8e-444b-aa16-dd8f9d2cb0ee', 'Bambang Tri', 'A', '2023-02-28 11:08:22.69835', '2023-02-28 11:19:48.458771', '', '', true);
INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('eb6a8af6-a218-4f41-a96e-6b982096bc89', 'Bening Gumilar', 'B', '2023-03-07 10:43:23.418977', NULL, '', NULL, false);


--
-- TOC entry 2871 (class 2606 OID 17027)
-- Name: jenis_mitra jenis_mitra_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jenis_mitra
    ADD CONSTRAINT jenis_mitra_pkey PRIMARY KEY (id);


--
-- TOC entry 2879 (class 2606 OID 17058)
-- Name: kelas_siswa_detail kelas_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kelas_siswa_detail
    ADD CONSTRAINT kelas_detail_pkey PRIMARY KEY (id);


--
-- TOC entry 2877 (class 2606 OID 17052)
-- Name: kelas_siswa kelas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kelas_siswa
    ADD CONSTRAINT kelas_pkey PRIMARY KEY (id);


--
-- TOC entry 2875 (class 2606 OID 17039)
-- Name: m_kelas m_kelas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_kelas
    ADD CONSTRAINT m_kelas_pkey PRIMARY KEY (id);


--
-- TOC entry 2873 (class 2606 OID 17033)
-- Name: m_siswa m_siswa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_siswa
    ADD CONSTRAINT m_siswa_pkey PRIMARY KEY (id);


--
-- TOC entry 3013 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-03-07 11:41:05 WIB

--
-- PostgreSQL database dump complete
--

