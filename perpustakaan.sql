PGDMP  ;                    }            perpustakaan    14.10    16.0 &               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    18582    perpustakaan    DATABASE     �   CREATE DATABASE perpustakaan WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE perpustakaan;
                postgres    false                        2615    2200    public    SCHEMA     2   -- *not* creating schema, since initdb creates it
 2   -- *not* dropping schema, since initdb creates it
                postgres    false                       0    0    SCHEMA public    ACL     Q   REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;
                   postgres    false    4            �            1259    18715    books    TABLE     %  CREATE TABLE public.books (
    id integer NOT NULL,
    title text NOT NULL,
    author text NOT NULL,
    published_year bigint,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);
    DROP TABLE public.books;
       public         heap    postgres    false    4            �            1259    18714    books_id_seq    SEQUENCE     �   CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.books_id_seq;
       public          postgres    false    214    4                       0    0    books_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;
          public          postgres    false    213            �            1259    18726    fine_settings    TABLE     �   CREATE TABLE public.fine_settings (
    id integer NOT NULL,
    fine_per_day numeric,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
 !   DROP TABLE public.fine_settings;
       public         heap    postgres    false    4            �            1259    18725    fine_settings_id_seq    SEQUENCE     �   CREATE SEQUENCE public.fine_settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.fine_settings_id_seq;
       public          postgres    false    4    216                       0    0    fine_settings_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.fine_settings_id_seq OWNED BY public.fine_settings.id;
          public          postgres    false    215            �            1259    18598    loans    TABLE     �  CREATE TABLE public.loans (
    id integer NOT NULL,
    user_id bigint NOT NULL,
    book_id bigint NOT NULL,
    loan_date timestamp without time zone DEFAULT now(),
    due_date timestamp without time zone NOT NULL,
    returned boolean DEFAULT false,
    return_date timestamp without time zone,
    deleted_at timestamp without time zone,
    borrow_date timestamp with time zone NOT NULL,
    fine numeric DEFAULT 0,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.loans;
       public         heap    postgres    false    4            �            1259    18597    loans_id_seq    SEQUENCE     �   CREATE SEQUENCE public.loans_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.loans_id_seq;
       public          postgres    false    212    4                       0    0    loans_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.loans_id_seq OWNED BY public.loans.id;
          public          postgres    false    211            �            1259    18584    users    TABLE     �  CREATE TABLE public.users (
    id integer NOT NULL,
    name text,
    email text,
    password text NOT NULL,
    role text,
    created_at timestamp with time zone,
    CONSTRAINT users_role_check CHECK ((role = ANY (ARRAY[('siswa'::character varying)::text, ('guru'::character varying)::text, ('admin'::character varying)::text, ('kepala_sekolah'::character varying)::text, ('pustakawan'::character varying)::text])))
);
    DROP TABLE public.users;
       public         heap    postgres    false    4            �            1259    18583    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    4    210                       0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    209            p           2604    18718    books id    DEFAULT     d   ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);
 7   ALTER TABLE public.books ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    213    214            s           2604    18729    fine_settings id    DEFAULT     t   ALTER TABLE ONLY public.fine_settings ALTER COLUMN id SET DEFAULT nextval('public.fine_settings_id_seq'::regclass);
 ?   ALTER TABLE public.fine_settings ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            l           2604    18601    loans id    DEFAULT     d   ALTER TABLE ONLY public.loans ALTER COLUMN id SET DEFAULT nextval('public.loans_id_seq'::regclass);
 7   ALTER TABLE public.loans ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    211    212    212            k           2604    18587    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    210    209    210                      0    18715    books 
   TABLE DATA           f   COPY public.books (id, title, author, published_year, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    214   $+                 0    18726    fine_settings 
   TABLE DATA           ]   COPY public.fine_settings (id, fine_per_day, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    216   �+                 0    18598    loans 
   TABLE DATA           �   COPY public.loans (id, user_id, book_id, loan_date, due_date, returned, return_date, deleted_at, borrow_date, fine, created_at, updated_at) FROM stdin;
    public          postgres    false    212   �+                 0    18584    users 
   TABLE DATA           L   COPY public.users (id, name, email, password, role, created_at) FROM stdin;
    public          postgres    false    210   i,                   0    0    books_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.books_id_seq', 1, true);
          public          postgres    false    213            !           0    0    fine_settings_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.fine_settings_id_seq', 1, true);
          public          postgres    false    215            "           0    0    loans_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.loans_id_seq', 3, true);
          public          postgres    false    211            #           0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 27, true);
          public          postgres    false    209            }           2606    18724    books books_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.books DROP CONSTRAINT books_pkey;
       public            postgres    false    214                       2606    18734     fine_settings fine_settings_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.fine_settings
    ADD CONSTRAINT fine_settings_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.fine_settings DROP CONSTRAINT fine_settings_pkey;
       public            postgres    false    216            {           2606    18605    loans loans_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.loans
    ADD CONSTRAINT loans_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.loans DROP CONSTRAINT loans_pkey;
       public            postgres    false    212            v           2606    18736    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public            postgres    false    210            x           2606    18593    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    210            �           1259    18790    idx_fine_settings_deleted_at    INDEX     \   CREATE INDEX idx_fine_settings_deleted_at ON public.fine_settings USING btree (deleted_at);
 0   DROP INDEX public.idx_fine_settings_deleted_at;
       public            postgres    false    216            y           1259    18770    idx_loans_deleted_at    INDEX     L   CREATE INDEX idx_loans_deleted_at ON public.loans USING btree (deleted_at);
 (   DROP INDEX public.idx_loans_deleted_at;
       public            postgres    false    212            �           2606    18754    loans loans_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.loans
    ADD CONSTRAINT loans_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 B   ALTER TABLE ONLY public.loans DROP CONSTRAINT loans_user_id_fkey;
       public          postgres    false    3192    210    212               j   x�3��HU(JL�M,�L��ӋssS�4BRKRS49�R*<J�J�\�RB2�s�9��@�����������������������1>�?�=... x� �         D   x�Eɱ� �:��>����<D&`�9 U5�� �
VZ!����>����{2�5��}�^Y*"	��         g   x����	� �g��n�<��T�;��C@�y�1�L�2
�^tr�L��2�A'��Μ�n�Z�A�Ƥ;�r�Ć����P4��U�F�d=8�.o�*Q         �  x�e�Kw�@F���`�nNz��lV��1*�(��ɦyA�����8Yd�Mթ�Ž�m˰P,Q����Џd��<C}*��cu�qj^=Z�d��rQa�	�t���xG����q��Y��/1*������ 5A3U��0z�$beU�2���t���Զ�Wi�erV�i:]W.�O���|a�G߀�����;k�m���|�W5S�q�3��WQ+�U���3�O��4�Dt׶o��q�!�9���M1�ɡ{�~���v����qJv�R�_�Ah
�&S%Xc�)��h��|ѕRdB��l�"8������ʺ��I|���y�D�Y]R��f'�;��S��Gmw�ƛ?y�u5��Sa���?
&0S�e����{��o>�T     